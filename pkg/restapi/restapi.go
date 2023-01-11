package restapi

import (
	"context"
	"net/http"

	"github.com/fidesy/go-url-shortener/pkg/database"
)

// BindAddr - server running port
// DBURL - database connection string
// DBName - name of database ('postgresql', ...)
type RestAPIConfig struct {
	BindAddr string
	DBURL    string
	DBName   string
}


type RestAPI struct {
	config *RestAPIConfig
	router *http.ServeMux
	db     database.Database
}

// RestAPI constructor
func New(config *RestAPIConfig) (*RestAPI, error) {
	db, err := database.NewDatabase(config.DBName)
	if err != nil {
		return nil, err
	}

	return &RestAPI{
		config: config,
		router: http.NewServeMux(),
		db:     db,
	}, nil
}

func (api *RestAPI) Start(ctx context.Context) error {
	api.configureRouters()
	if err := api.configureDatabase(ctx); err != nil {
		return err
	}
	defer api.db.Close()

	return http.ListenAndServe(api.config.BindAddr, api.router)
}

func (api *RestAPI) configureDatabase(ctx context.Context) error {
	err := api.db.Open(ctx, api.config.DBURL)
	return err
}
