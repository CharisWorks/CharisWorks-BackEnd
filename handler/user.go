package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForUser(firebaseApp validation.IFirebaseApp, UserRequests users.IRequests) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(firebaseMiddleware(firebaseApp))
	UserRouter.Use(userMiddleware(UserRequests))
	{
		UserRouter.GET("/user", func(ctx *gin.Context) {
			userId := ctx.GetString("userId")
			User, err := UserRequests.UserGet(userId)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(http.StatusOK, User)
		})
		UserRouter.DELETE("/user", func(ctx *gin.Context) {
			userId := ctx.GetString("userId")
			err := UserRequests.UserDelete(userId)
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
			err = UserRequests.UserProfileUpdate(userId, *profile)
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
			err = UserRequests.UserAddressRegister(userId, *payload)
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
			err = UserRequests.UserAddressUpdate(userId, *payload)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
		})
	}
}
