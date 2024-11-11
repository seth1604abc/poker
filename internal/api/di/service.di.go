package di

import (
	"poker/internal/api/service"

	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	service.NewUserService,
)