//go:build wireinject
// +build wireinject

package main

import (
	"adx-admin/api"
	"adx-admin/internal/domian/service"
	"adx-admin/internal/infrastructure/persistence"
	"adx-admin/pkg/configer"
	"adx-admin/pkg/core/app"
	"adx-admin/pkg/database"
	"adx-admin/pkg/microlog"
	"adx-admin/pkg/server"
	"github.com/google/wire"
)

func wireApp(*configer.Config, microlog.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		database.ProviderSet,
		persistence.ProviderSet,
		service.ProviderSet,
		api.ProviderSet,
		server.Provider,
		server.NewServer,
		app.NewApp,
	))
}
