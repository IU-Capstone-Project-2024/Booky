package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Ip   string
	Port string
}

type S3Config struct {
	Endpoint string
	Bucket   string
}

type StorageConfig struct {
	CourseStorage string
	NoteStorage   string
	FileStorage   string
	UserStorage   string

	S3 S3Config
}

type GptConfig struct {
	NoteImprovementPrompt string

	YandexGPT YandexGPTConfig
}

type YandexGPTConfig struct {
	URL      string
	ModelUri string
	ApiKey   string
}

type Config struct {
	AppEnv  string
	Server  ServerConfig
	Storage StorageConfig
	Gpt     GptConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	courseStorage := flag.String("courseStorage", "in-memory", "Course storage option (in-memory, PostgreSQL)")
	noteStorage := flag.String("noteStorage", "in-memory", "Note storage option (in-memory, PostgreSQL)")
	fileStorage := flag.String("fileStorage", "in-memory", "File storage option (in-memory, PostgreSQL, S3)")
	userStorage := flag.String("userStorage", "in-memory", "User storage option (in-memory, PostgreSQL)")

	flag.Parse()

	config := &Config{
		AppEnv: getEnv("APP_ENV", "production"),
		Server: ServerConfig{
			Ip:   getEnv("BOOKY_API_IP", "0.0.0.0"),
			Port: getEnv("BOOKY_API_PORT", "4000"),
		},
		Storage: StorageConfig{
			CourseStorage: *courseStorage,
			NoteStorage:   *noteStorage,
			FileStorage:   *fileStorage,
			UserStorage:   *userStorage,
			S3: S3Config{
				Endpoint: getEnv("BOOKY_S3_ENDPOINT", ""),
				Bucket:   getEnv("BOOKY_S3_BUCKET", ""),
			},
		},
		Gpt: GptConfig{
			NoteImprovementPrompt: getEnv("BOOKY_NOTE_IMPROVEMENT_PROMPT", ""),
			YandexGPT: YandexGPTConfig{
				URL:      getEnv("YANDEX_GPT_URL", ""),
				ModelUri: getEnv("YANDEX_GPT_MODEL_URI", ""),
				ApiKey:   getEnv("YANDEX_GPT_API_KEY", ""),
			},
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
