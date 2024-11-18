package main

import (
	"cisco-catalyst-prometheus-exporter/configHandler"
	"cisco-catalyst-prometheus-exporter/Logger"
	"log/slog"
	"net/http"
)

var DEFAULT_PORT string = "9000"

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

func startTLSService(lg *slog.Logger, cfg map[string]interface{}) {
	port := DEFAULT_PORT
	if cfg_port, exist := cfg["server"].(map[string]interface{})["port"]; exist {
		port = cfg_port.(string)
	}
	var crt, key string = "", ""
	if cfg_certificate, exist := cfg["certificate"].(map[string]interface{}); exist {
    crt, _ = cfg_certificate["crt"].(string)
		key, _ = cfg_certificate["key"].(string) 
	} else {
		panic("TLS certicifates missing in configuration!")
	}
  lg.Info("Starting https server at port " + port)
	http.HandleFunc("/metrics", metricsHandler)
	if err := http.ListenAndServeTLS(":" + port, crt, key, nil); err != nil {
		lg.Error("ListenAndServeTLS: " + err.Error())
	}
}

func startService(lg *slog.Logger, cfg map[string]interface{}) {
	port := DEFAULT_PORT
	if cfg_port, exist := cfg["server"].(map[string]interface{})["port"]; exist {
		port = cfg_port.(string)
	}
	lg.Info("Starting http service at port " + port)
	http.HandleFunc("/metrics", metricsHandler)
  if err := http.ListenAndServe(":" +  port, nil); err != nil {
		lg.Error("ListenAndServe: " + err.Error())
	}
}

func main() {
	lg := Logger.InitLogger()
	var config map[string]interface{}
	configHandler.GetConfig(lg, &config)
	
	if c, exist := config["certificate"]; exist{ 
		if _, exist := c.(map[string]interface{})["crt"]; exist {
	    startTLSService(lg, config)
    } else {
	  	startService(lg, config)
	  }
  } else {
		startService(lg, config)
	}
}