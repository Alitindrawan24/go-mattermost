package user

import "github.com/gin-gonic/gin"

func (h *Handler) HandleGetUser(ctx *gin.Context) {
	user, err := h.user.GetUser(ctx)
	if err != nil {
		ctx.JSON(400, err.Error())
	}

	ctx.JSON(200, user)
}
