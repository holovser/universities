//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bachelor-service/app"
	"github.com/bachelor-service/handler"
	school "github.com/bachelor-service/service"
	"github.com/google/wire"
)

func NewApp() (app.App, error) {
	panic(wire.Build(
		wire.Struct(new(app.App), "*"),
		ProviderHandler,
		wire.Bind(new(handler.IHandler), new(*handler.Handler)),
		ProviderService,
		wire.Bind(new(school.IService), new(*school.Service)),
		ProviderDb,
	))

	return app.App{}, nil
}
