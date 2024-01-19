package handler

import (
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router *gin.Engine
}

func NewHandler(router *gin.Engine) *Handler {
	return &Handler{
		Router: router,
	}
}

func (h *Handler) SetupRoutes(firebaseApp *validation.FirebaseApp) {
	firebase := h.Router.Group("/firebase")
	{
		firebase.GET("/test", func(c *gin.Context) {
			app := validation.NewFirebaseApp()
			app.VerifyIDToken(c, "[idToken]")
		})
	}
}
