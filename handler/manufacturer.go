package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForManufacturer(firebaseApp validation.IFirebaseApp, manufacturerRequests manufacturer.IItemRequests, userRequests users.IRequests) {
	UserRouter := h.Router.Group("/api/products")
	UserRouter.Use(firebaseMiddleware(firebaseApp))
	{
		UserRouter.Use((userMiddleware(userRequests)))
		UserRouter.Use(manufacturerMiddleware())
		{
			UserRouter.POST("/", func(ctx *gin.Context) {
				payload, err := utils.GetPayloadFromBody(ctx, &manufacturer.RegisterPayload{})
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				userId := ctx.GetString("userId")
				err = manufacturerRequests.Register(payload, userId, utils.GenerateRandomString())
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly registered")
			})
			UserRouter.PATCH("/:item_id", func(ctx *gin.Context) {
				payload, err := utils.GetPayloadFromBody(ctx, &manufacturer.UpdatePayload{})
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				userId := ctx.GetString("userId")
				itemId, err := utils.GetParams("item_id", ctx)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				err = manufacturerRequests.Update(payload, userId, *itemId)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly updated")
			})
			UserRouter.DELETE("/:item_id", func(ctx *gin.Context) {
				itemId, err := utils.GetParams("item_id", ctx)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				userId := ctx.GetString(string(userId))
				err = manufacturerRequests.Delete(*itemId, userId)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly deleted")
			})
		}

	}
}
