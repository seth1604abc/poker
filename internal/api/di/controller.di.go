package di

import (
	"poker/internal/api/controller"

	"github.com/google/wire"
)

var ControllerSet = wire.NewSet(
	controller.NewUserController,
)