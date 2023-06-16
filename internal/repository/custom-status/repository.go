package customstatus

import (
	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	"github.com/gin-gonic/gin"
)

type RepositoryInterface interface {
	UpdateCustomStatus(ctx *gin.Context, request entity.CustomStatusRequest) error
}

type Repository struct {
}

func New() RepositoryInterface {
	return &Repository{}
}
