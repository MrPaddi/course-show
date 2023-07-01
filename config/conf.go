package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	SqliteDBPath string
	LoggingPath  string
	Port         string
}

func New(configFile string) *Config {
	// configFile if not exist or not readable, Fatal
	if _, err := os.Stat(configFile); os.IsNotExist(err) || os.IsPermission(err) {
		log.Fatal("Config file not found")
	}
	// read config file,if error, Fatal
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	// instance config by file(json)
	config := &Config{}
	// decode file(json) to config, if error, Fatal
	if err := json.NewDecoder(file).Decode(config); err != nil {
		log.Fatal(err)
	}
	// return config
	return config
}
