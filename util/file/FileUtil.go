package file

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	TestQueue string `json:"test_queue"`
}

var config = &Configuration{}

const defaultPath = "conf.json"

func init() {
	//path can be read from SYSTEM PATH
	configPath := defaultPath
	if envPath := os.Getenv("config"); envPath != "" {
		configPath = envPath
	}
	f, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Can't read config from %v , reson: %v", configPath, err.Error())
	}
	json.Unmarshal(f, config)
}

func GetConf() *Configuration {
	return config
}
