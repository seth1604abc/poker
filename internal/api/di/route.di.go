package di

import (
	"poker/internal/api/route"

	"github.com/google/wire"
)

var RouteSet = wire.NewSet(
	route.NewUserRoute,
)