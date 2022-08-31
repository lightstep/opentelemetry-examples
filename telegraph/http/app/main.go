package main

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

type HeapStat struct {
	Name  string `json:"name"`
	Value uint64 `json:"value"`
	Time  string `json:"time"`
}

func main() {
	http.HandleFunc("/heapbasics", func(w http.ResponseWriter, r *http.Request) {
		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)

		heapMetrics := []*HeapStat{
			{
				Name:  "idle",
				Value: ms.HeapIdle,
				Time:  time.Now().Format("unix"),
			},
			{
				Name:  "inuse",
				Value: ms.HeapInuse,
				Time:  time.Now().Format("unix"),
			},
			{
				Name:  "reserved",
				Value: ms.HeapSys,
				Time:  time.Now().Format("unix"),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(heapMetrics)
	})

	http.ListenAndServe(":8080", nil)
}
