package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp validation.IFirebaseApp, CartRequests cart.ICartRequests, CartDB cart.ICartDB, UserRequests users.IUserRequests, CartUtils cart.ICartUtils, UserDB users.IUserDB) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(firebaseMiddleware(firebaseApp))
	{
		CartRouter.Use(userMiddleware(UserRequests, UserDB))
		{
			CartRouter.GET("/", func(ctx *gin.Context) {
				Cart, err := CartRequests.Get(ctx, CartDB, CartUtils)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, Cart)
			})
			CartRouter.POST("/", func(ctx *gin.Context) {

				err := CartRequests.Register(CartDB, CartUtils, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully registered")
			})
			CartRouter.DELETE("/", func(ctx *gin.Context) {
				err := CartRequests.Delete(CartDB, CartUtils, ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully deleted")
			})
		}
	}
}
