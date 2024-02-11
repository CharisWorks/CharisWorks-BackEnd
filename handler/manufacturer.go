package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForManufacturer(firebaseApp validation.IFirebaseApp) {
	UserRouter := h.Router.Group("/api/products")
	UserRouter.Use(firebaseMiddleware(&validation.FirebaseApp{}))
	{
		UserRouter.Use(manufacturerMiddleware())
		{
			UserRouter.POST("", func(ctx *gin.Context) {
				bindBody := new(manufacturer.ItemRegisterPayload)
				payload, err := getPayloadFromBody(ctx, &bindBody)
				if err != nil {
					return
				}
				err = manufacturer.RegisterItem(**payload, manufacturer.ExampleManufacturerRequests{}, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly registered")
			})
			UserRouter.PATCH("", func(ctx *gin.Context) {
				bindBody := new(manufacturer.ItemUpdatePayload)
				payload, err := getPayloadFromBody(ctx, &bindBody)
				if err != nil {
					return
				}
				err = manufacturer.UpdateItem(**payload, manufacturer.ExampleManufacturerRequests{}, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly updated")
			})
			UserRouter.DELETE("", func(ctx *gin.Context) {
				itemId := ctx.Query("item_id")
				err := manufacturer.DeleteItem(itemId, manufacturer.ExampleManufacturerRequests{}, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly deleted")
			})
		}

	}
}
