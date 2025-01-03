package main

import (
	"log"

	application "github.com/arakhimiy/edu-connect/internal/app"
	"github.com/arakhimiy/edu-connect/internal/config"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	app := application.NewApp(cfg)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
