package handler

import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/authstatus"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForAuthStatus() {
	h.Router.POST("/api/userauthstatus", func(c *gin.Context) {
		// レスポンスの処理
		i := new(struct{ email string })
		if err := c.BindJSON(&i); err != nil {
			log.Print(err)
		}
		PreviewList := authstatus.AuthStatusCheck(i.email, authstatus.AuthStatusRequests{})
		c.JSON(200, PreviewList)
	})

}
