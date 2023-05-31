package main

import(
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	// "golang.org/x/exp/slices"
)

type Record struct {
	Name string
	Description string
	Unit string
	DataType string
	Attributes []string
}

func parseLogsFile(inputFile string) ([]Record, error) {
	 in, err := os.Open(inputFile)
	 if err != nil {
		 return nil, fmt.Errorf("failed to open input file: %w", err)
	 }
	 defer in.Close()

	 scanner := bufio.NewScanner(in)
	 records := make([]Record, 0)
	 record := Record{}
	 isAttribute := false
	 // records := make(map[string]struct{}, 0)
	 for scanner.Scan() {
		 line := strings.TrimSpace(scanner.Text())
		 if strings.Contains(line, "->") {
			 split := strings.Split(line, "->")
			 if len(split) == 2 {
				 value := strings.TrimSpace(split[1])
				 split = strings.Split(value, ":") 
				 if len(split) == 2 {
					 value = strings.TrimSpace(split[1])
				 }
				 switch {
				 case strings.Contains(line, "Name:"):
					 record.Name = value
				 case strings.Contains(line, "Description:"):
					 record.Description = value
				 case strings.Contains(line, "Unit:"):
					 record.Unit = value
				case strings.Contains(line, "DataType:"):
					record.DataType = value
				default:
					if isAttribute {
						record.Attributes = append(record.Attributes, value)
					}
				}
			}
		} else if strings.Contains(line, "Data point attributes:") {
			isAttribute = true
		} else if strings.Contains(line, "Metric #") {
			if record.Name != "" {
				record = Record{}
			}
			isAttribute = false
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input file: %w", err)
	}
	// slices.SortFunc(records), func(a, b Record) { return a.Name < b.Name })

	return records, nil
}

func writeCSV(outputFile string, records []Record) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"Name", "Description ", "Unit", "DataType", "Attributes"}); err != nil {
		return err
	}

	for _, record := range records {
		if err := writer.Write([]string{record.Name, record.Description, record.Unit, record.DataType, strings.Join(record.Attributes, ";")}); err != nil {
			return err
		}
	}

	return nil
}

func orderMetricsForDisplay(records []Record) []Record {
	// make a map keyed on the 6 char prefixes of metric names 
	prefixMap := make(map[string][]Record)
	for _, v := range records {
		if len(v.Name) > 5 {
			prefixMap[v.Name[:6]] = append(prefixMap[v.Name[:6]], v)
		} else {
			prefixMap[v.Name] = append(prefixMap[v.Name], v)
		}
	}

	// output the records in order (1) 6 char prefix of .Name and (2) order in logs
	var groupedRecords []Record
	for len(prefixMap) > 0 {
		var largest int
		for _, v := range prefixMap {
			if len(v) > largest {
				largest = len(v)
			}
		}

		for k, v := range prefixMap {
			if len(v) == largest {
				groupedRecords = append(groupedRecords, v...)
				delete(prefixMap, k)
			}
		}
	}
	return groupedRecords
}

func processDirectories(rootDir string) error {
	entries, err := ioutil.ReadDir(rootDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDirPath := filepath.Join(rootDir, entry.Name())
			logFilePath := filepath.Join(subDirPath, "logs.txt")
			outputFilePath := filepath.Join(subDirPath, "metrics.csv")

			if _, err := os.Stat(logFilePath); !os.IsNotExist(err) {
				records, err := parseLogsFile(logFilePath)
				if err != nil {
					return fmt.Errorf("failed to process file: '%s': %w", logFilePath, err)
				}

				orderedRecords := orderMetricsForDisplay(records)

				if err := writeCSV(outputFilePath, orderedRecords); err != nil {
					return fmt.Errorf("failed to write output file '%s': %w", outputFilePath, err)
				}
			}

			if err := processDirectories(subDirPath); err != nil {
				return err
			}
		}
	}
	return nil
}


func main() {
	if err := processDirectories("collector"); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

