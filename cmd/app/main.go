package main

import (
	"log"
	"ntheanh201-journal/config"
	"ntheanh201-journal/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
