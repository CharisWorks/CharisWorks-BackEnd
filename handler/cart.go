package handler

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp *validation.FirebaseApp) {
	CartRouter := h.Router.Group("/api/cart")
	//CartRouter.Use(firebaseMiddleware(*firebaseApp))
	{
		CartRouter.GET("", func(c *gin.Context) {
			Cart := cart.GetCart(cart.CartRequest{})
			c.JSON(http.StatusOK, Cart)
		})
		CartRouter.POST("", func(c *gin.Context) {
			bindBody := new(cart.CartRequestPayload)
			payload, err := getPayloadFromBody(c, &bindBody)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			log.Print("Payload: ", (*payload))
			log.Println(payload)
			Cart := cart.PostCart(**payload, cart.CartRequest{})
			c.JSON(http.StatusOK, Cart)
		})
		CartRouter.PATCH("", func(c *gin.Context) {
			bindBody := new(cart.CartRequestPayload)
			payload, err := getPayloadFromBody(c, &bindBody)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
			Cart := cart.UpdateCart(**payload, cart.CartRequest{})
			c.JSON(http.StatusOK, Cart)
		})
		CartRouter.DELETE("", func(c *gin.Context) {
			ItemId := c.Query("item_id")
			Cart := cart.DeleteCart(ItemId, cart.CartRequest{})
			c.JSON(http.StatusOK, Cart)
		})
	}
}
