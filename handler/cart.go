package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp validation.IFirebaseApp, CartRequests cart.ICartRequests, CartDB cart.ICartRepository, UserRequests users.IUserRequests, CartUtils cart.ICartUtils, UserDB users.IUserDB) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(firebaseMiddleware(firebaseApp))
	{
		CartRouter.Use(userMiddleware(UserRequests, UserDB))
		{
			CartRouter.GET("/", func(ctx *gin.Context) {
				userId := ctx.GetString("userId")
				Cart, err := CartRequests.Get(userId, CartDB, CartUtils)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, Cart)
			})
			CartRouter.POST("/", func(ctx *gin.Context) {
				userId := ctx.GetString("userId")
				cartRequestPayload, err := utils.GetPayloadFromBody(ctx, &cart.CartRequestPayload{})
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				err = CartRequests.Register(userId, *cartRequestPayload, CartDB, CartUtils)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully registered")
			})
			CartRouter.DELETE("/", func(ctx *gin.Context) {
				userId := ctx.GetString("userId")
				itemId, err := utils.GetQuery("itemId", ctx)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				err = CartRequests.Delete(userId, *itemId, CartDB, CartUtils)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully deleted")
			})
		}
	}
}
