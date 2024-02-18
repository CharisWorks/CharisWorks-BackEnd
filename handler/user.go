package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForUser(firebaseApp validation.IFirebaseApp, i user.IUserRequests) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(firebaseMiddleware(firebaseApp))
	{
		UserRouter.GET("/user", func(ctx *gin.Context) {
			User, err := i.UserGet(ctx.MustGet("UserId").(string), ctx)
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, User)
		})
		UserRouter.DELETE("/user", func(ctx *gin.Context) {
			err := i.UserDelete(ctx.MustGet("UserId").(string), ctx)
			if err != nil {
				return
			}
		})
		UserRouter.POST("/profile", func(ctx *gin.Context) {
			bindBody := new(user.UserProfileRegisterPayload)
			payload, err := getPayloadFromBody(ctx, &bindBody)
			if err != nil {
				return
			}
			err = i.UserProfileRegister(**payload, ctx)
			if err != nil {
				return
			}
		})
		UserRouter.PATCH("/profile", func(ctx *gin.Context) {
			bindBody := new(user.UserProfile)
			payload, err := getPayloadFromBody(ctx, &bindBody)
			if err != nil {
				return
			}
			err = i.UserProfileUpdate(**payload, ctx)
			if err != nil {
				return
			}
		})
		UserRouter.POST("/address", func(ctx *gin.Context) {
			bindBody := new(user.UserAddressRegisterPayload)
			payload, err := getPayloadFromBody(ctx, &bindBody)
			if err != nil {
				return
			}
			err = i.UserAddressRegister(**payload, ctx)
			if err != nil {
				return
			}
		})
		UserRouter.PATCH("/address", func(ctx *gin.Context) {
			bindBody := new(user.UserAddress)
			payload, err := getPayloadFromBody(ctx, &bindBody)
			if err != nil {
				return
			}
			err = i.UserAddressUpdate(**payload, ctx)
			if err != nil {
				return
			}
		})
	}
}
