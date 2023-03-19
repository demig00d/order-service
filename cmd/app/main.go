package main

import (
	"log"

	"github.com/demig00d/order-service/config"
	"github.com/demig00d/order-service/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
