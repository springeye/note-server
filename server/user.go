package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/login", func(writer http.ResponseWriter, request *http.Request) {

	})
	r.Post("/register", func(writer http.ResponseWriter, request *http.Request) {

	})
	return r
}
