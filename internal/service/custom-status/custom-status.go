package customstatus

import (
	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) SetCustomStatus(ctx *gin.Context, request entity.CustomStatusRequest) error {
	err := s.customstatus.UpdateCustomStatus(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
