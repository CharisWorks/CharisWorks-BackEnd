package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp *validation.FirebaseApp) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(firebaseMiddleware(*firebaseApp))
	{
		CartRouter.GET("", func(c *gin.Context) {
			Cart, err := cart.GetCart(cart.CartRequest{})
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			c.JSON(http.StatusOK, Cart)
		})
		CartRouter.POST("", func(c *gin.Context) {
			bindBody := new(cart.CartRequestPayload)
			payload, err := getPayloadFromBody(c, &bindBody)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			err = cart.PostCart(**payload, cart.CartRequest{})
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			c.JSON(http.StatusOK, "Item was successfully registered")
		})
		CartRouter.PATCH("", func(c *gin.Context) {
			bindBody := new(cart.CartRequestPayload)
			payload, err := getPayloadFromBody(c, &bindBody)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			err = cart.UpdateCart(**payload, cart.CartRequest{})
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			c.JSON(http.StatusOK, "Item was successfully updated")
		})
		CartRouter.DELETE("", func(c *gin.Context) {
			itemId := c.Query("item_id")
			if itemId == "" {
				c.JSON(http.StatusBadRequest, "cannot get itemId")
				return
			}
			err := cart.DeleteCart(itemId, cart.CartRequest{})
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			c.JSON(http.StatusOK, "Item was successfully deleted")
		})
	}
}
