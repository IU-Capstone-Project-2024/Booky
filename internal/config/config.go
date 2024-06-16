package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AppEnv string `yaml:"app_env" env:"APP_ENV"`
	Ip     string `yaml:"ip" env:"IP"`
	Port   string `yaml:"port" env:"PORT"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		AppEnv: getEnv("APP_ENV", "production"),
		Ip:     getEnv("IP", "127.0.0.1"),
		Port:   getEnv("PORT", "8080"),
	}

	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
