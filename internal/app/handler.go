package app

import (
	customstatus "github.com/Alitindrawan24/go-mattermost/internal/handler/custom-status"
	"github.com/Alitindrawan24/go-mattermost/internal/handler/user"
)

type Handlers struct {
	User         *user.Handler
	CustomStatus *customstatus.Handler
}

func NewHandler(service *Services) *Handlers {
	return &Handlers{
		User:         user.New(service.user),
		CustomStatus: customstatus.New(service.customstatus),
	}
}
