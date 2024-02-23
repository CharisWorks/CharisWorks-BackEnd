package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForManufacturer(firebaseApp validation.IFirebaseApp, ManufacturerRequests manufacturer.IManufacturerRequests) {
	UserRouter := h.Router.Group("/api/products")
	UserRouter.Use(firebaseMiddleware(firebaseApp))
	{
		UserRouter.Use(manufacturerMiddleware())
		{
			UserRouter.POST("/", func(ctx *gin.Context) {
				bindBody := new(manufacturer.ItemRegisterPayload)
				payload, err := getPayloadFromBody(ctx, &bindBody)
				if err != nil {
					return
				}
				err = ManufacturerRequests.RegisterItem(**payload, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly registered")
			})
			UserRouter.PATCH("/", func(ctx *gin.Context) {
				bindBody := new(manufacturer.ItemUpdatePayload)
				payload, err := getPayloadFromBody(ctx, &bindBody)
				if err != nil {
					return
				}
				err = ManufacturerRequests.UpdateItem(**payload, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly updated")
			})
			UserRouter.DELETE("/", func(ctx *gin.Context) {
				itemId, err := getQuery("item_id", true, ctx)
				if err != nil {
					return
				}
				err = ManufacturerRequests.DeleteItem(*itemId, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly deleted")
			})
		}

	}
}
