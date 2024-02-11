package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/authstatus"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForAuthStatus() {
	h.Router.POST("/api/userauthstatus", func(c *gin.Context) {
		// レスポンスの処理
		bind := new(authstatus.Email)
		payload, err := getPayloadFromBody(c, &bind)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		PreviewList, err := authstatus.AuthStatusCheck(**payload, authstatus.ExampleAuthStatusRequests{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(200, PreviewList)
	})

}
