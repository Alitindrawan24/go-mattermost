package user

import (
	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	"github.com/Alitindrawan24/go-mattermost/internal/repository/user"
	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	GetUser(ctx *gin.Context) (entity.User, error)
}

type Service struct {
	user user.RepositoryInterface
}

func New(user user.RepositoryInterface) ServiceInterface {
	return &Service{user}
}
