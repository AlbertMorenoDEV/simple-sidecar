package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func accessLoggerMiddleware(next http.Handler) http.Handler {
	return http.Handler(handlers.LoggingHandler(os.Stdout, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})))
}
