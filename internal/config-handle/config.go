package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/reenphygeorge/light-server/internal/logger"
)

// Config structure
type Config struct {
	StartPort int `json:"startPort"`
	SkipDirectories []string `json:"skipDirectories"`
	RootPath string `json:"rootPath"`
}

// Get values from config and insert it to struct object
func GetValues(configObject *Config) {
	filePath := "lsconfig.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error()
	}
	err = json.Unmarshal(data, configObject)
	if err != nil {
		log.Fatal("Error parsing config file:", err)
	}
}