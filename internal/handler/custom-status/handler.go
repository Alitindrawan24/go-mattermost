package customstatus

import customstatus "github.com/Alitindrawan24/go-mattermost/internal/service/custom-status"

type Handler struct {
	customstatus customstatus.ServiceInterface
}

func New(customstatus customstatus.ServiceInterface) *Handler {
	return &Handler{customstatus}
}
