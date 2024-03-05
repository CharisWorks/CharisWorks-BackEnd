package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp validation.IFirebaseApp, cartRequests cart.IRequests, cartRepository cart.IRepository, userRequests users.IRequests, cartUtils cart.IUtils, userRepository users.IRepository) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(firebaseMiddleware(firebaseApp))
	{
		CartRouter.Use(userMiddleware(userRequests, userRepository))
		{
			CartRouter.GET("/", func(ctx *gin.Context) {
				userId := ctx.GetString("userId")
				Cart, err := cartRequests.Get(userId, cartRepository, cartUtils)
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
				err = cartRequests.Register(userId, *cartRequestPayload, cartRepository, cartUtils)
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
				err = cartRequests.Delete(userId, *itemId, cartRepository, cartUtils)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully deleted")
			})
		}
	}
}
