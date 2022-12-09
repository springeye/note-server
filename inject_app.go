package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/springeye/oplin/server"
)

func providerApplication(mainRouter chi.Router) *application {
	return &application{
		mainRouter: mainRouter,
	}
}
func providerRouter() chi.Router {
	return server.MainRouter()
}
