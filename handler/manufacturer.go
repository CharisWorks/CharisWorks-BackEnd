package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForManufacturer(firebaseApp *validation.FirebaseApp) {
	UserRouter := h.Router.Group("/api/products")
	UserRouter.Use(firebaseMiddleware(*firebaseApp))
	{
		UserRouter.Use(manufacturerMiddleware(*firebaseApp))
		{
			UserRouter.POST("", func(c *gin.Context) {
				bindBody := new(manufacturer.ItemRegisterPayload)
				payload, err := getPayloadFromBody(c, &bindBody)
				if err != nil {
					c.JSON(http.StatusBadRequest, err)
					return
				}
				err = manufacturer.RegisterItem(**payload, manufacturer.ExampleManufacturerRequests{})
				if err != nil {
					c.JSON(http.StatusBadRequest, err)
					return
				}
				c.JSON(http.StatusOK, "Item was successfuly registered")
			})
			UserRouter.PATCH("", func(c *gin.Context) {
				bindBody := new(manufacturer.ItemUpdatePayload)
				payload, err := getPayloadFromBody(c, &bindBody)
				if err != nil {
					c.JSON(http.StatusBadRequest, err)
					return
				}
				err = manufacturer.UpdateItem(**payload, manufacturer.ExampleManufacturerRequests{})
				if err != nil {
					c.JSON(http.StatusBadRequest, err)
					c.Abort()
				}
				c.JSON(http.StatusOK, "Item was successfuly updated")
			})
			UserRouter.DELETE("", func(c *gin.Context) {
				itemId := c.Query("item_id")
				err := manufacturer.DeleteItem(itemId, manufacturer.ExampleManufacturerRequests{})
				if err != nil {
					c.JSON(http.StatusBadRequest, err)
					c.Abort()
				}
				c.JSON(http.StatusOK, "Item was successfuly deleted")
			})
		}

	}
}
