//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitApplication() *application {
	panic(wire.Build(
		providerAppConfig,
		providerStore,
		providerCommand,
		providerServer,
		providerApplication))
}
