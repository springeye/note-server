//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitApplication() *application {
	panic(wire.Build(providerRouter,providerAppConfig, providerApplication))
}
