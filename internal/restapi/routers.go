package restapi

import (
	"context"
	"fmt"
	"github.com/fidesy/go-url-shortener/internal/mw"
	"net/http"
	"strings"
)

func (api *RestAPI) configureRouters() {
	// /create?url=https://someurl.com
	api.router.Handle("/create", mw.Logger(api.createURL))
	api.router.Handle("/", mw.Logger(api.redirect))
}

func (api *RestAPI) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	hash := strings.ReplaceAll(r.URL.Path, "/", "")
	if hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	originalURL, err := api.db.GetOriginalURL(context.Background(), hash)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusPermanentRedirect)
}

func (api *RestAPI) createURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{'error': 'url query param is missing'}"))
		return
	}

	shortURL, err := api.db.CreateShortURL(context.Background(), url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s:%s/%s", api.config.Host, api.config.Port, shortURL)))
}
