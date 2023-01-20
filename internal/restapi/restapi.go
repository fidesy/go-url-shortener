package restapi

import (
	"context"
	"github.com/fidesy/go-url-shortener/internal/databases/database"
	"log"
	"net/http"
)

// RestAPIConfig
// BindAddr - server running port
// DBURL - database connection string
// DBName - name of database ('postgresql', ...)
type RestAPIConfig struct {
	Host   string
	Port   string
	DBURL  string
	DBName string
}

type RestAPI struct {
	config *RestAPIConfig
	router *http.ServeMux
	db     database.Database
}

// New RestAPI constructor
func New(config *RestAPIConfig) (*RestAPI, error) {
	api := &RestAPI{}

	api.config = config

	db, err := database.New(config.DBName)
	if err != nil {
		return nil, err
	}

	api.db = db

	api.router = http.NewServeMux()

	return api, nil
}

func (api *RestAPI) Start(ctx context.Context) error {
	api.configureRouters()
	if err := api.configureDatabase(ctx); err != nil {
		return err
	}
	defer api.db.Close()

	log.Println("Server started")
	return http.ListenAndServe(":"+api.config.Port, api.router)
}

func (api *RestAPI) configureDatabase(ctx context.Context) error {
	err := api.db.Open(ctx, api.config.DBURL)
	return err
}
