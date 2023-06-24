package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/reenphygeorge/servette/internal/logger"
)

// Config structure
type Config struct {
	Port            int      `json:"port"`
	SkipDirectories []string `json:"skipDirectories"`
}

// Get values from config and insert it to struct object
func GetValues(configObject *Config) {
	filePath := "srt.config.json"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		configObject.Port = 5500
		configObject.SkipDirectories = append(configObject.SkipDirectories, "git")
	} else {
		data, err := ioutil.ReadFile(filePath)
		err = json.Unmarshal(data, configObject)
		if err != nil {
			logger.Error("")
		}
	}
}
