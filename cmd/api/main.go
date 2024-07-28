package main

import (
	"log"

	"github.com/mchayapol/go-task-app/config"
	"github.com/mchayapol/go-task-app/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Printf("Config initialized")

	app := server.NewApp()

	log.Printf("Starting server on port %s...", viper.GetString("port"))

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
