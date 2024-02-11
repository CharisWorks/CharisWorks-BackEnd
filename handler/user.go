package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForUser(firebaseApp validation.IFirebaseApp) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(firebaseMiddleware(firebaseApp))
	{
		UserRouter.GET("/user", func(ctx *gin.Context) {
			User, err := user.UserGet(ctx.MustGet("UserId").(string), user.ExampleUserRequests{}, ctx)
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, User)
		})
		UserRouter.DELETE("/user", func(ctx *gin.Context) {

		})
		UserRouter.POST("/profile", func(ctx *gin.Context) {

		})
		UserRouter.PATCH("/profile", func(ctx *gin.Context) {

		})
		UserRouter.POST("/address", func(ctx *gin.Context) {})
		UserRouter.PATCH("/address", func(ctx *gin.Context) {})
	}
}
