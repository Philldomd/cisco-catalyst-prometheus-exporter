package configHandler

import (
	"fmt"
	"os"
	//"strings"
	//"strconv"

	"gopkg.in/yaml.v3"
)

func check(e error) {
  if e != nil {
		panic(e)
	}
}

func loadConfigFromFile(path string, c *map[string]interface{}) {
	if _, err := os.Stat(path); err != nil {
		fmt.Println("ERROR: File not found!")
		return
	} else {
	  f, err := os.ReadFile(path)
    check(err)
		err = yaml.Unmarshal(f, c)
	  check(err)
	}
}

func assignEnvironmentValues(c *map[string]interface{}) {
	/*for _, e := range os.Environ(){
		switch value := strings.Split(e, "="); value[0] {
		case "SERVER_NAME":
			c.Server.Name = value[1]
		case "SERVER_PORT":
			c.Server.Port,_ = strconv.Atoi(value[1])
		case "CERT_CRT":
			c.Certificate.Crt = value[1]
		case "CERT_KEY":
			c.Certificate.Key = value[1]
		}
	}*/
}

/*Tries to load config from config files. If ENV values is present this will
have higher priority. Default path: /var/cisco-dna/config.yaml
*/
func GetConfig(c *map[string]interface{}) {
	if os.Getenv("DNA_CONFIG_PATH") != ""{
	  loadConfigFromFile(os.Getenv("DNA_CONFIG_PATH"), c)
	} else {
		loadConfigFromFile("/var/cisco-dna/config.yaml", c)
	}
	assignEnvironmentValues(c)
}