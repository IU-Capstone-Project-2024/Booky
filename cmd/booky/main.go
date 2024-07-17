package main

import (
	"booky-back/internal/config"
	"booky-back/internal/pkg/logger"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}

	app := NewApp(config)

	err = app.Run()
	if err != nil {
		logger.Fatal(err)
	}
}
