package CDNALogger

import (
	"log/slog"
	"os"
	"strings"
)

var CDNA_LOG_LEVEL slog.Leveler
var CDNA_LOG_PATH *os.File

func configLogger() {
	for _, e := range os.Environ() {
		value := strings.Split(e, "=")
		if value[0] == "CDNA_LOG_LEVEL" {
			switch value[1] {
		    case "Debug":
		      CDNA_LOG_LEVEL = slog.LevelDebug
				case "Warning":
					CDNA_LOG_LEVEL = slog.LevelWarn
				case "Error":
					CDNA_LOG_LEVEL = slog.LevelError
				default:
					CDNA_LOG_LEVEL = slog.LevelInfo
		  } 
		} else if value[0] == "CDNA_LOG_PATH" {
		  if strings.Contains(value[1], "/") {
				file, err := os.Create(value[1])
				if err != nil {
					panic("Could not open file for writing!")
				}
				CDNA_LOG_PATH = file
			} else {
				CDNA_LOG_PATH = os.Stdout
			}
		}
	}
}

func InitLogger() (*slog.Logger){
	configLogger()
	th := slog.NewTextHandler(CDNA_LOG_PATH, &slog.HandlerOptions{ Level: CDNA_LOG_LEVEL })
	return slog.New(th)
}