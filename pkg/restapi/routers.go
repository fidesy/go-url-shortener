package restapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func (api *RestAPI) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return 
	}

	hash := strings.ReplaceAll(r.URL.Path, "/", "")
	if hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	original_url, err := api.db.GetOriginalURL(context.Background(), hash)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, original_url, http.StatusPermanentRedirect)
}

func (api *RestAPI) createURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return 
	}

	query := r.URL.Query()
	url := query.Get("url")

	if url == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	short_url, err := api.db.CreateURL(context.Background(), url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("http://localhost%s/%s\n", api.config.BindAddr, short_url)))
}
