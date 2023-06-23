package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/reenphygeorge/servette/internal/logger"
)

// Config structure
type Config struct {
	Port            int      `json:"port"`
	SkipDirectories []string `json:"skipDirectories"`
	RootPath        string   `json:"rootPath"`
}

// Get values from config and insert it to struct object
func GetValues(configObject *Config) {
	filePath := "srv.config.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error("Config file not found! ")
		os.Exit(0)
	}
	err = json.Unmarshal(data, configObject)
	if err != nil {
		log.Fatal("Error parsing config file:", err)
	}
}
