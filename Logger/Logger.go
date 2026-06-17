package Logger

import (
	"log/slog"
	"os"
	"strings"
)

var CCPE_LOG_LEVEL slog.Leveler
var CCPE_LOG_PATH *os.File = os.Stdout

func configLogger() {
	for _, e := range os.Environ() {
		value := strings.Split(e, "=")
		switch value[0] {
			case "CCPE_LOG_LEVEL":
				switch value[1] {
					case "Debug":
						CCPE_LOG_LEVEL = slog.LevelDebug
					case "Warning":
						CCPE_LOG_LEVEL = slog.LevelWarn
					case "Error":
						CCPE_LOG_LEVEL = slog.LevelError
					default:
						CCPE_LOG_LEVEL = slog.LevelInfo
				}
			case "CCPE_LOG_PATH":
			if strings.Contains(value[1], "/") {
				file, err := os.Create(value[1])
				if err != nil {
					panic("Could not open file for writing!")
				}
				CCPE_LOG_PATH = file
			}
		}
	}
}

func InitLogger() (*slog.Logger){
	configLogger()
	th := slog.NewTextHandler(CCPE_LOG_PATH, &slog.HandlerOptions{ Level: CCPE_LOG_LEVEL })
	return slog.New(th)
}