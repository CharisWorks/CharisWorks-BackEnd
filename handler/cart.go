package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp validation.IFirebaseApp) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(firebaseMiddleware(firebaseApp))
	{
		CartRouter.GET("", func(ctx *gin.Context) {
			Cart, err := cart.GetCart(cart.CartRequest{}, ctx)
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, Cart)
		})
		CartRouter.POST("", func(ctx *gin.Context) {
			bindBody := new(cart.CartRequestPayload)
			payload, err := getPayloadFromBody(ctx, &bindBody)
			if err != nil {
				return
			}
			err = cart.PostCart(**payload, cart.CartRequest{}, ctx)
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, "Item was successfully registered")
		})
		CartRouter.PATCH("", func(ctx *gin.Context) {
			bindBody := new(cart.CartRequestPayload)
			payload, err := getPayloadFromBody(ctx, &bindBody)
			if err != nil {
				return
			}
			err = cart.UpdateCart(**payload, cart.CartRequest{}, ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, "Item was successfully updated")
		})
		CartRouter.DELETE("", func(ctx *gin.Context) {
			itemId := ctx.Query("item_id")
			if itemId == "" {
				ctx.JSON(http.StatusBadRequest, "cannot get itemId")
				return
			}
			err := cart.DeleteCart(itemId, cart.CartRequest{}, ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, "Item was successfully deleted")
		})
	}
}
