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
		BindAddr: os.Getenv("BIND_ADDR"),
		DBURL:    os.Getenv("DBURL"),
		DBName:   os.Getenv("DBName"),
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
