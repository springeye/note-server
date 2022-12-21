package main

import (
	"github.com/springeye/oplin/cmd"
	"github.com/springeye/oplin/config"
	"github.com/springeye/oplin/db"
	"github.com/springeye/oplin/server"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"os"
)

type application struct {
	conf   *config.AppConfig
	store  *db.Store
	cmd    *cmd.Command
	server *server.Server
}
type command struct {
	db *gorm.DB
}

func (receiver *application) start() error {
	port := receiver.conf.Port
	if port <= 0 {
		port = 3000
	}
	return receiver.server.Start(port)
}
func (receiver *application) init() error {
	loggerOpts := slog.HandlerOptions{
		AddSource: true,
	}
	if receiver.conf.Debug {
		loggerOpts.Level = slog.DebugLevel
	} else {
		loggerOpts.Level = slog.ErrorLevel
	}
	slog.SetDefault(slog.New(loggerOpts.NewTextHandler(os.Stdout)))

	slog.Debug("init database")
	receiver.store.Setup()
	return nil
}
