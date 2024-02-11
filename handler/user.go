package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForUser(firebaseApp *validation.FirebaseApp) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(firebaseMiddleware(*firebaseApp))
	{
		UserRouter.GET("/user", func(c *gin.Context) {
			User, err := user.UserGet(c.MustGet("UserId").(string), user.ExampleUserRequests{})
			if err != nil {
				c.JSON(http.StatusNotFound, err)
				c.Abort()
			}
			c.JSON(http.StatusOK, User)
		})
	}
}
