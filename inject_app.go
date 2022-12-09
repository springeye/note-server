package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"github.com/springeye/oplin/config"
	"github.com/springeye/oplin/server"
)

func providerApplication(mainRouter chi.Router,conf *config.AppConfig) *application {
	return &application{
		mainRouter: mainRouter,
		conf: conf,
	}
}
func providerRouter() chi.Router {
	return server.MainRouter()
}
func providerAppConfig() *config.AppConfig {
	config.Setup("config.json")
	var conf config.AppConfig
	viper.Unmarshal(conf)
	return &conf
}
