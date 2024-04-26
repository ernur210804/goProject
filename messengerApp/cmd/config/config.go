package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database *DatabaseConfig
	Server   *ServerConfig
	// Other configurations
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type ServerConfig struct {
	Port string
	// Other server configuration parameters
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("cmd/config") // Adjust the path to your config directory
	viper.SetConfigType("yaml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error loading config file: %s", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config file: %s", err))
	}

	return &config
}
