package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/user"
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

func firebaseMiddleware(app validation.FirebaseApp) gin.HandlerFunc {
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
func manufacturerMiddleware(app validation.FirebaseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		User := user.UserGet(c.MustGet("UserId").(string), user.UserRequests{})
		if User == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Firebase Account exists but this account was not found in Backend Server"})
			c.Abort()
		}
		if User == nil || User.Manufacturer == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Account is not manufacturer"})
			c.Abort()
		}

		//内部の実行タイミング
		c.Next()

	}
}
