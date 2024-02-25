package handler

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForUser(firebaseApp validation.IFirebaseApp, UserRequests user.IUserRequests, UserDB user.IUserDB) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(firebaseMiddleware(firebaseApp))
	UserRouter.Use(userMiddleware(UserRequests, UserDB))
	{
		UserRouter.GET("/user", func(ctx *gin.Context) {
			User, err := UserRequests.UserGet(ctx, UserDB)
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, User)
		})
		UserRouter.DELETE("/user", func(ctx *gin.Context) {
			err := UserRequests.UserDelete(ctx, UserDB)
			log.Print(err)
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User was successfully deleted"})
		})

		UserRouter.PATCH("/profile", func(ctx *gin.Context) {

			err := UserRequests.UserProfileUpdate(ctx, UserDB)
			if err != nil {
				return
			}
		})
		UserRouter.POST("/address", func(ctx *gin.Context) {

			err := UserRequests.UserAddressRegister(ctx, UserDB)
			if err != nil {
				return
			}
		})
		UserRouter.PATCH("/address", func(ctx *gin.Context) {

			err := UserRequests.UserAddressUpdate(ctx, UserDB)
			if err != nil {
				return
			}
		})
	}
}
