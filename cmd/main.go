package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/fidesy/go-url-shortener/pkg/restapi"
)
 
func main() {
	err := godotenv.Load()
	checkError(err)

	api := restapi.New(&restapi.RestAPIConfig{
		BindAddr:      os.Getenv("BIND_ADDR"),
		DBURI:         os.Getenv("DBURI"),
	})

	err = api.Start(context.Background())
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
