package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Ip   string
	Port string
}

type GptConfig struct {
	Token                 string
	NoteImprovementPrompt string
}

type Config struct {
	AppEnv string
	Server ServerConfig
	Gpt    GptConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		AppEnv: getEnv("APP_ENV", "production"),
		Server: ServerConfig{
			Ip:   getEnv("BOOKY_API_IP", "0.0.0.0"),
			Port: getEnv("BOOKY_API_PORT", "4000"),
		},
		Gpt: GptConfig{
			Token:                 getEnv("BOOKY_GPT_TOKEN", ""),
			NoteImprovementPrompt: getEnv("BOOKY_NOTE_IMPROVEMENT_PROMPT", ""),
		},
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
