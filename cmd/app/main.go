package main

import (
	"log"

	"github.com/okankaraduman/goFinalApp/config"
	"github.com/okankaraduman/goFinalApp/internal/app"
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
