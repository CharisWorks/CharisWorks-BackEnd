package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForUser(firebaseApp validation.IFirebaseApp, userRequests users.IRequests, userRepository users.IRepository, userUtils users.IUtils) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(firebaseMiddleware(firebaseApp))
	UserRouter.Use(userMiddleware(userRequests, userRepository))
	{
		UserRouter.GET("/user", func(ctx *gin.Context) {
			userId := ctx.GetString("userId")
			User, err := userRequests.UserGet(userId, userRepository)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(http.StatusOK, User)
		})
		UserRouter.DELETE("/user", func(ctx *gin.Context) {
			userId := ctx.GetString("userId")
			err := userRequests.UserDelete(userId, userRepository)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User was successfully deleted"})
		})

		UserRouter.PATCH("/profile", func(ctx *gin.Context) {
			profile, err := utils.GetPayloadFromBody(ctx, &users.UserProfile{})
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			userId := ctx.GetString("userId")
			err = userRequests.UserProfileUpdate(userId, *profile, userRepository, userUtils)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
		})
		UserRouter.POST("/address", func(ctx *gin.Context) {
			payload, err := utils.GetPayloadFromBody(ctx, &users.UserAddressRegisterPayload{})
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			userId := ctx.GetString("userId")
			err = userRequests.UserAddressRegister(userId, *payload, userRepository, userUtils)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
		})
		UserRouter.PATCH("/address", func(ctx *gin.Context) {
			payload, err := utils.GetPayloadFromBody(ctx, &users.UserAddress{})
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			userId := ctx.GetString("userId")
			err = userRequests.UserAddressUpdate(userId, *payload, userRepository, userUtils)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
		})
	}
}
