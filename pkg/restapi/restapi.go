package restapi

import (
	"context"
	"net/http"

	"github.com/fidesy/go-url-shortener/pkg/database"
)

type RestAPIConfig struct {
	BindAddr string
	DBURL    string
}

type RestAPI struct {
	config *RestAPIConfig
	router *http.ServeMux
	db     database.Database
}

func New(config *RestAPIConfig) *RestAPI {
	return &RestAPI{
		config: config,
		router: http.NewServeMux(),
		db:     database.NewPostgreSQL(),
	}
}

func (api *RestAPI) Start(ctx context.Context) error {
	api.configureRouters()
	if err := api.configureDatabase(ctx); err != nil {
		return err
	}
	defer api.db.Close(context.Background())

	return http.ListenAndServe(api.config.BindAddr, api.router)
}

func (api *RestAPI) configureDatabase(ctx context.Context) error {
	err := api.db.Open(ctx, api.config.DBURL)
	return err
}
