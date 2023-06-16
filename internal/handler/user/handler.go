package user

import "github.com/Alitindrawan24/go-mattermost/internal/service/user"

type Handler struct {
	user user.ServiceInterface
}

func New(user user.ServiceInterface) *Handler {
	return &Handler{user}
}
