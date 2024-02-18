package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp validation.IFirebaseApp) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(firebaseMiddleware(firebaseApp))
	{
		CartRouter.Use(userMiddleware(user.ExampleUserRequests{}))
		{
			CartRouter.GET("", func(ctx *gin.Context) {
				Cart, err := cart.GetCart(cart.ExapleCartRequest{}, ctx)
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
				err = cart.PostCart(**payload, cart.ExapleCartRequest{}, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully registered")
			})
			CartRouter.DELETE("", func(ctx *gin.Context) {
				itemId, err := getQuery("item_id", ctx)
				if err != nil {
					return
				}
				err = cart.DeleteCart(*itemId, cart.ExapleCartRequest{}, ctx)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, err)
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully deleted")
			})
		}
	}
}
