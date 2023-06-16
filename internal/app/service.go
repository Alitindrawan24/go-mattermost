package app

import (
	customstatus "github.com/Alitindrawan24/go-mattermost/internal/service/custom-status"
	"github.com/Alitindrawan24/go-mattermost/internal/service/user"
)

type Services struct {
	user         user.ServiceInterface
	customstatus customstatus.ServiceInterface
}

func NewService(repository *Repositories) *Services {
	return &Services{
		user:         user.New(repository.user),
		customstatus: customstatus.New(repository.customstatus),
	}
}
