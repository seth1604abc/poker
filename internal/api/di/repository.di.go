package di

import (
	"poker/internal/api/repository"

	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	repository.NewUserRepository,
)