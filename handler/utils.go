package handler

import (
	"net/http"

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
			idToken := "[idToken]"
			app, err := validation.NewFirebaseApp()
			if err != nil {
				return
			}
			app.VerifyIDToken(c, idToken)
		})
	}
}

func FirebaseMiddleware(app validation.FirebaseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		UID, err := app.VerifyIDToken(c, c.Request.Header.Get("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
		}
		c.Set("UserId", UID)
		//内部の実行タイミング
		c.Next()

	}
}
