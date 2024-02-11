package handler

import (
	"github.com/charisworks/charisworks-backend/internal/authstatus"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForAuthStatus() {
	h.Router.POST("/api/userauthstatus", func(ctx *gin.Context) {
		// レスポンスの処理
		bind := new(authstatus.Email)
		payload, err := getPayloadFromBody(ctx, &bind)
		if err != nil {
			return
		}
		PreviewList, err := authstatus.AuthStatusCheck(**payload, authstatus.ExampleAuthStatusRequests{}, ctx)
		if err != nil {
			return
		}
		ctx.JSON(200, PreviewList)
	})

}
