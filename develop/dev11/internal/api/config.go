package api

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration settings for the application,
// including the address and port to which the application should bind,
// specified by the 'addr_port' field in the YAML configuration file.
type Config struct {
	AddrPort string `yaml:"addr_port"`
}

// NewConfig initializes a new Config object by attempting to load configuration from a specified YAML file.
// If the file does not exist or an error occurs during reading or parsing, it returns a default configuration.
func NewConfig(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Println("Config file not found. Loading default config.")
		return &Config{":8080"}
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("Error reading YAML file. Loading default config.")
		return &Config{":8080"}
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Println("Error parsing YAML file. Loading default config.")
		return &Config{":8080"}
	}

	return &config
}
