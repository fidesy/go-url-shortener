package restapi

import (
	"context"
	"net/http"

	"github.com/fidesy/go-url-shortener/pkg/postgresdb"
)

type RestAPIConfig struct {
	BindAddr string
	DBURI    string
}

type RestAPI struct {
	config *RestAPIConfig
	router *http.ServeMux
	db     *postgresdb.PostgresDB
}

func New(config *RestAPIConfig) *RestAPI {
	return &RestAPI{
		config: config,
		router: http.NewServeMux(),
		db:     postgresdb.New(),
	}
}

func (api *RestAPI) Start(ctx context.Context) error {
	api.configureRouters()
	if err := api.configureDatabase(ctx); err != nil {
		return err
	}
	defer api.db.Close()

	return http.ListenAndServe(api.config.BindAddr, api.router)
}

func (api *RestAPI) configureRouters() {
	// /create?url=https://someurl.com
	api.router.HandleFunc("/create", api.createURL)
	api.router.HandleFunc("/", api.redirect)
}

func (api *RestAPI) configureDatabase(ctx context.Context) error {
	err := api.db.Open(ctx, api.config.DBURI)
	return err
}
