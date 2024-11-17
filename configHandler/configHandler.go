package configHandler

import (
	"log/slog"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var logger *slog.Logger

func check(e error) {
  if e != nil {
		logger.Error(e.Error())
		panic(e)
	}
}

func loadConfigFromFile(path string, c *map[string]interface{}) {
	if _, err := os.Stat(path); err != nil {
		logger.Error("File not found!")
		return
	} else {
	  f, err := os.ReadFile(path)
    check(err)
		err = yaml.Unmarshal(f, c)
	  check(err)
	}
}

func configTemplate(c *map[string]interface{}, rootKey string, key string, value string) {
  if m, ok := (*c)[rootKey].(map[string]interface{}); ok {
		logger.Warn(rootKey + "." + key + ": " + m[key].(string) + " overwritten by environment variable!: " + value)
		m[key] = value
	} else {
		(*c)[rootKey] = map[string]interface{}{ key: value }
	}
}

func assignEnvironmentValues(c *map[string]interface{}) {
	for _, e := range os.Environ(){
		switch value := strings.Split(e, "="); value[0] {
		  case "CDNA_SERVER_NAME",
		      "CDNA_SERVER_PORT",
		      "CDNA_CERTIFICATE_CRT",
		      "CDNA_CERTIFICATE_KEY",
		      "CDNA_CISCO_DNA_URL",
		      "CDNA_CISCO_TOKEN":
		    logger.Debug(e)
		    keys := strings.SplitN(value[0], "_", 2)
			  logger.Debug("Reading environment variable: " + strings.ToLower(keys[0] + " " + keys[1]) + " = " + value[1])
			  configTemplate(c, strings.ToLower(keys[0]), strings.ToLower(keys[1]), value[1])
		}
	}
}

/*Tries to load config from config files. If ENV values is present this will
have higher priority. Default path: /var/cisco-dna/config.yaml
*/
func GetConfig(lg *slog.Logger, c *map[string]interface{}) {
	logger = lg
	if os.Getenv("CDNA_CONFIG_PATH") != ""{
	  loadConfigFromFile(os.Getenv("DNA_CONFIG_PATH"), c)
	} else {
		loadConfigFromFile("/var/cisco-dna/config.yaml", c)
	}
	assignEnvironmentValues(c)
}