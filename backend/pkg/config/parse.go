package config

import (
	"fmt"
	"os"

	"github.com/jesseduffield/yaml"
)

var config *Config

func ParseConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

func GetConfig() *Config {
	if config == nil {
		panic("config not initialized")
	}
	return config
}
