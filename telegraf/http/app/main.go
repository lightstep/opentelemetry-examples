package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	hostname    = "unset"
	metricsPath = "/heapbasics"
)

func init() {
	h, err := os.Hostname()
	if err != nil {
		hostname = "undetected"
	}
	hostname = h
}

type Fields struct {
	Idle     uint64 `json:"idle"`
	Inuse    uint64 `json:"inuse"`
	Reserved uint64 `json:"reserved"`
}

type Tags struct {
	Host        string `json:"hostname"`
	MetricsPath string `json:"metrics_path"`
}

type HeapStats struct {
	Name      string `json:"name"`
	Fields    Fields `json:"fields"`
	Tags      Tags   `json:"tags"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	http.HandleFunc(metricsPath, func(w http.ResponseWriter, r *http.Request) {
		log.Println("metrics requested")
		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)

		heapMetrics := HeapStats{
			Name: "heapstats",
			Fields: Fields{
				Idle:     ms.HeapIdle,
				Inuse:    ms.HeapInuse,
				Reserved: ms.HeapSys,
			},
			Tags: Tags{
				Host:        hostname,
				MetricsPath: metricsPath,
			},
			Timestamp: time.Now().Unix(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(heapMetrics)
		log.Println("metrics provided")
	})

	http.ListenAndServe(":8080", nil)
}
