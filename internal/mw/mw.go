package mw

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s | %s | %s\n", req.Method, req.URL.RequestURI(), time.Since(start))
	})
}
