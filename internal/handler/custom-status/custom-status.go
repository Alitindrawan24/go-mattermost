package customstatus

import (
	"net/http"

	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleSetCustomStatus(ctx *gin.Context) {

	var customStatusRequest entity.CustomStatusRequest
	if err := ctx.ShouldBindJSON(&customStatusRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customStatusRequest.Emoji = "calendar"

	err := h.customstatus.SetCustomStatus(ctx, customStatusRequest)
	if err != nil {
		ctx.JSON(400, err.Error())
	}

	ctx.JSON(200, "Successfully set custom status")
}
