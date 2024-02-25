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

				err := ManufacturerRequests.RegisterItem(ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly registered")
			})
			UserRouter.PATCH("/", func(ctx *gin.Context) {

				err := ManufacturerRequests.UpdateItem(ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly updated")
			})
			UserRouter.DELETE("/", func(ctx *gin.Context) {

				err := ManufacturerRequests.DeleteItem(ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfuly deleted")
			})
		}

	}
}
