package customstatus

import (
	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	customstatus "github.com/Alitindrawan24/go-mattermost/internal/repository/custom-status"
	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	SetCustomStatus(ctx *gin.Context, request entity.CustomStatusRequest) error
}

type Service struct {
	customstatus customstatus.RepositoryInterface
}

func New(customstatus customstatus.RepositoryInterface) ServiceInterface {
	return &Service{customstatus}
}
