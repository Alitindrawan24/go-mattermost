package user

import (
	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetUser(ctx *gin.Context) (entity.User, error) {
	user, err := s.user.GetUser(ctx)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
