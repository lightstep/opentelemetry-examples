package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	// "time"
	"os"

	metricService "github.com/lightstep/metric_language_equivalent_tests/internal/opentelemetry-proto-gen/collector/metrics/v1"
	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
)

type (
	testServer struct {
		data []byte
	}
)

var (
	ErrUnsupported = fmt.Errorf("unsupported method")
)

func (t *testServer) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL", r.URL)
	fmt.Println("HEADERS", r.Header)

	defer r.Body.Close()
	data, _ := ioutil.ReadAll(r.Body)

	fmt.Println("BODY", string(data))

	w.Header().Set("Content-Type", "application/json")
	w.Write(t.data)
}

func main() {

	the_testServer := testServer{}

	go func() {
		listener, err := net.Listen("tcp", "0.0.0.0:7001")
		if err != nil {
			log.Fatal(err)
			fmt.Println("err")
		}
		grpcServer := grpc.NewServer()
		metricService.RegisterMetricsServiceServer(grpcServer, &the_testServer)
		go grpcServer.Serve(listener)
		defer grpcServer.Stop()

		select {}
	}()

	go func() {
		http.HandleFunc("/", the_testServer.handler)
		log.Fatal(http.ListenAndServe(":7002", nil))
	}()

	fmt.Fprintln(os.Stderr, "Server ready")
	select {}
}

func (t *testServer) Export(ctx context.Context, req *metricService.ExportMetricsServiceRequest) (*metricService.ExportMetricsServiceResponse, error) {
	var emptyValue = metricService.ExportMetricsServiceResponse{}
	_, ok := grpcMetadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("failed")
	}
	data, _ := json.MarshalIndent(req, "", "  ")
	fmt.Println(string(data))

	t.data = data
	fmt.Fprintln(os.Stderr, "data ready")
	return &emptyValue, nil
}
