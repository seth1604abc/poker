//go:build wireinject
// +build wireinject

package di

import (
	"poker/internal/api/route"

	"github.com/google/wire"
)

func InitializeApp() (*route.Route, error) {
	wire.Build(
		DBSet,
		RepositorySet,
		ServiceSet,
		ControllerSet,
		RouteSet,
		route.NewRoute,
	)

	return &route.Route{}, nil
}