package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"go-echo-rest-api/router"
)

type M map[string]interface{}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.New()

	if err := r.Start(":9000"); err != nil {
		log.Fatal("Error run server")
	}
}
