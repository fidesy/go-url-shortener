package main

import (
	"context"
	"github.com/fidesy/go-url-shortener/internal/restapi"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	checkError(err)

	api, err := restapi.New(&restapi.RestAPIConfig{
		Host:   os.Getenv("HOST"),
		Port:   os.Getenv("PORT"),
		DBURL:  os.Getenv("DB_URL"),
		DBName: os.Getenv("DB_NAME"),
	})
	checkError(err)

	err = api.Start(context.Background())
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
