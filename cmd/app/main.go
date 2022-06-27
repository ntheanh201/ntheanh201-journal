package main

import (
	"github.com/joho/godotenv"
	"log"
	"ntheanh201-journal/config"
	"ntheanh201-journal/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	app.Run(cfg)
}
