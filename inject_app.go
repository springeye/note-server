package main

import (
	"github.com/spf13/viper"
	"github.com/springeye/oplin/cmd"
	"github.com/springeye/oplin/config"
	"github.com/springeye/oplin/db"
	"github.com/springeye/oplin/server"
)

func providerApplication(conf *config.AppConfig, store *db.Store, c *cmd.Command, s *server.Server) *application {
	return &application{
		conf:   conf,
		store:  store,
		cmd:    c,
		server: s,
	}
}
func providerAppConfig() *config.AppConfig {
	config.Setup("config.json")
	var conf config.AppConfig
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
func providerStore(appConfig *config.AppConfig) *db.Store {
	dbStore := db.Store{Conf: appConfig}
	return &dbStore
}
func providerCommand(store *db.Store) *cmd.Command {
	return cmd.NewCommand(store)
}
func providerServer(store *db.Store) *server.Server {
	return server.NewServer(store)
}
