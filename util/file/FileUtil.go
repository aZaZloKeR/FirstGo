package file

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	testQueue string
}

func GetConf(name string) string {
	f, err := os.ReadFile("conf.json")
	if err != nil {
		log.Println(err)
	}
	var data map[string]string
	json.Unmarshal(f, &data)

	for k, v := range data {
		if k == name {
			return v
		}
	}
	return ""
}
