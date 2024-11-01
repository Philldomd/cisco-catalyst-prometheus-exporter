package main

import (
	"log"
	"net/http"
	"cisco-dna-prometheus-exporter/configHandler"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/metrics" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusBadRequest)
		return
	}
	if r.Header.Get("Accept") == "application/openmetrics-text" {
	  w.Header().Set("Content-Type", "application/openmetrics-text")
	} else {
    w.Header().Set("Content-Type", "text/plain")
	}
	w.Write([]byte("Hello!\n"))
}

func main() {
	config = configHandler()
	lg := log.Default()
	lg.Printf("Starting server at port 9000\n")
	http.HandleFunc("/metrics", metricsHandler)
	if err := http.ListenAndServeTLS(":9000", "certificate/server.crt", "certificate/server.key", nil); err != nil {
		lg.Fatal("ListenAndServe: ", err)
	}
}
