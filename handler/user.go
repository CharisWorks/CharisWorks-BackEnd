package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForUser(firebaseApp *validation.FirebaseApp) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(FirebaseMiddleware(*firebaseApp))
	{
		UserRouter.GET("/user", func(c *gin.Context) {
			User := user.ExampleUser(c.MustGet("UID").(string))
			c.JSON(http.StatusOK, User)
		})
	}
}
