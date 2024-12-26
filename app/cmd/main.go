package main

import (
	"log"

	application "gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/app"
	"gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/config"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	app := application.NewApp(cfg)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
