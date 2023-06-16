package user

import (
	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	"github.com/gin-gonic/gin"
)

type RepositoryInterface interface {
	GetUser(ctx *gin.Context) (entity.User, error)
}

type Repository struct {
}

func New() RepositoryInterface {
	return &Repository{}
}
