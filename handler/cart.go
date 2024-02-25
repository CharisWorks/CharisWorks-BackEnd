package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp validation.IFirebaseApp, CartRequests cart.ICartRequests, CartDB cart.ICartDB, UserRequests user.IUserRequests, CartUtils cart.ICartUtils, UserDB user.IUserDB) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(firebaseMiddleware(firebaseApp))
	{
		CartRouter.Use(userMiddleware(UserRequests, UserDB))
		{
			CartRouter.GET("/", func(ctx *gin.Context) {
				Cart, err := CartRequests.Get(ctx, CartDB, CartUtils, ctx.MustGet("UserId").(string))
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, Cart)
			})
			CartRouter.POST("/", func(ctx *gin.Context) {
				bindBody := new(cart.CartRequestPayload)
				payload, err := getPayloadFromBody(ctx, &bindBody)
				if err != nil {
					return
				}
				err = CartRequests.Register(**payload, CartDB, CartUtils, ctx, ctx.MustGet("UserId").(string))
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully registered")
			})
			CartRouter.DELETE("/", func(ctx *gin.Context) {
				itemId, err := getQuery("item_id", true, ctx)
				if err != nil {
					return
				}
				err = CartRequests.Delete(*itemId, CartDB, CartUtils, ctx, ctx.MustGet("UserId").(string))
				if err != nil {
					ctx.JSON(http.StatusBadRequest, err)
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully deleted")
			})
		}
	}
}
