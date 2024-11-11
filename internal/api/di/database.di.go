package di

import (
	"poker/internal/api/model"

	"github.com/google/wire"
)

var DBSet = wire.NewSet(
	model.InitMysql,
)