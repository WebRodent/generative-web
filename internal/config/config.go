package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	ConnStr string `yaml:"connStr"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func Load() (Config, error) {
	// Load and parse the configuration from a file or environment variables
	// Return the loaded configuration
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return Config{}, fmt.Errorf("Failed to read YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to parse YAML: %v", err)
	}

	return config, nil
}

func (cfg Config) Print() {
	fmt.Println("Server Configuration:")
	fmt.Println("Host:", cfg.Server.Host)
	fmt.Println("Port:", cfg.Server.Port)
	fmt.Println("Database Configuration:")
	fmt.Println("Connection String:", cfg.Database.ConnStr)
}
