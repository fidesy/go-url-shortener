package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/fidesy/go-url-shortener/internal/handler"
	"github.com/fidesy/go-url-shortener/internal/repository"
	"github.com/fidesy/go-url-shortener/internal/service"
	"github.com/fidesy/go-url-shortener/pkg/utils"
)

func main() {
	conf, err := utils.LoadConfig("./configs/config.yaml")
	checkError(err)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	defer cancel()

	repos, err := repository.NewRepository(ctx, conf)
	checkError(err)

	services := service.NewService(conf, repos)
	handlers := handler.NewHandler(services)
	routers := handlers.InitRoutes()

	go func() {
		err = http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), routers)
		checkError(err)
	}()

	<-ctx.Done()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
