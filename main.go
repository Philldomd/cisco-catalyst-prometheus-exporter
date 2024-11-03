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

func startTLSService(lg *log.Logger, cfg map[string]interface{}) {
	port := 9000
	if cfg_port, exist := cfg["server"].(map[string]interface{})["port"]; exist {
		port = cfg_port.(int)
	}
	var crt, key string = "", ""
	if cfg_certificate, exist := cfg["certificate"].(map[string]interface{}); exist {
    crt, _ = cfg_certificate["crt"].(string)
		key, _ = cfg_certificate["key"].(string) 
	} else {
		panic("TLS certicifates missing in configuration!")
	}
  lg.Printf("Starting https server at port %d\n", port)
	http.HandleFunc("/metrics", metricsHandler)
	if err := http.ListenAndServeTLS(":"+strconv.Itoa(port), crt, key, nil); err != nil {
		lg.Fatal("ListenAndServeTLS: ", err)
	}
}

func startService(lg *log.Logger, cfg map[string]interface{}) {
	port := 9000
	if cfg_port, exist := cfg["server"].(map[string]interface{})["port"]; exist {
		port = cfg_port.(int)
	}
	lg.Printf("Starting http service at port %d", port)
	http.HandleFunc("/metrics", metricsHandler)
  if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		lg.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	var config map[string]interface{}
	configHandler.GetConfig(&config)
	
	lg := log.Default()
	if _, exist := config["certificate"].(map[string]interface{})["crt"]; exist {
	  startTLSService(lg, config)
  } else {
		startService(lg, config)
	}
}