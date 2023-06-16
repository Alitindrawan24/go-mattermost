package app

import (
	customstatus "github.com/Alitindrawan24/go-mattermost/internal/repository/custom-status"
	"github.com/Alitindrawan24/go-mattermost/internal/repository/user"
)

type Repositories struct {
	user         user.RepositoryInterface
	customstatus customstatus.RepositoryInterface
}

func NewRepository() *Repositories {
	return &Repositories{
		user:         user.New(),
		customstatus: customstatus.New(),
	}
}
