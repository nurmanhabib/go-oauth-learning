package main

import (
	"gitlab.com/nurmanhabib/go-oauth2/interfaces/routes"
	"log"
	"os"
)

func main() {
	r := routes.New().Init()

	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		appPort = "8080"
	}

	log.Fatal(r.Run(":" + appPort))
}
