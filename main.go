package main

import (
	"cisco-dna-prometheus-exporter/configHandler"
	"log"
	"net/http"
	"strconv"
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

func startTLSService(lg *log.Logger, cfg configHandler.Config) {
  lg.Printf("Starting https server at port %d\n", cfg.Server.Port)
	http.HandleFunc("/metrics", metricsHandler)
	if err := http.ListenAndServeTLS(":"+strconv.Itoa(cfg.Server.Port), cfg.Certificate.Crt, cfg.Certificate.Key, nil); err != nil {
		lg.Fatal("ListenAndServeTLS: ", err)
	}
}

func startService(lg *log.Logger, cfg configHandler.Config) {
	lg.Printf("Starting http service at port %d", cfg.Server.Port)
	http.HandleFunc("/metrics", metricsHandler)
  if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), nil); err != nil {
		lg.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	var config configHandler.Config
	configHandler.GetConfig(&config)
	lg := log.Default()
	if config.Certificate.Crt != "" && config.Certificate.Key != "" {
	  startTLSService(lg, config)
  } else {
		startService(lg, config)
	}
}