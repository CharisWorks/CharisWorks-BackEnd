package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/cart"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp *validation.FirebaseApp) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(FirebaseMiddleware(*firebaseApp))
	{
		CartRouter.GET("", func(c *gin.Context) {
			Cart := cart.GetCart(cart.CartRequest{})
			c.JSON(http.StatusOK, Cart)
		})
		CartRouter.POST("", func(c *gin.Context) {
			CartRequestPayload := new(cart.CartRequestPayload)
			err := c.Bind(&CartRequestPayload)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			Cart := cart.PostCart(*CartRequestPayload, cart.CartRequest{})
			c.JSON(http.StatusOK, Cart)
		})
		CartRouter.PATCH("", func(c *gin.Context) {
			CartRequestPayload := new(cart.CartRequestPayload)
			err := c.Bind(&CartRequestPayload)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			Cart := cart.UpdateCart(*CartRequestPayload, cart.CartRequest{})
			c.JSON(http.StatusOK, Cart)
		})
		CartRouter.DELETE("", func(c *gin.Context) {
			ItemId := c.Query("item_id")
			Cart := cart.DeleteCart(ItemId, cart.CartRequest{})
			c.JSON(http.StatusOK, Cart)
		})
	}
}
