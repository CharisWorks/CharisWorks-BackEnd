package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForManufacturer(firebaseApp *validation.FirebaseApp) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(firebaseMiddleware(*firebaseApp))
	{
		UserRouter.Use(manufacturerMiddleware(*firebaseApp))
		{
			UserRouter.POST("/products", func(c *gin.Context) {
				ReqPayload := new(items.ItemOverviewProperties)
				err := c.BindJSON(&ReqPayload)
				if err != nil {
					c.JSON(http.StatusBadRequest, "Invalid Data")
					c.Abort()
				}
				err = manufacturer.RegisterItem(*ReqPayload, manufacturer.ExampleManufacturerRequests{})
				if err != nil {
					c.JSON(http.StatusBadRequest, err)
					c.Abort()
				}
				c.JSON(http.StatusOK, "Item was successfuly registered")
			})
		}

	}
}
