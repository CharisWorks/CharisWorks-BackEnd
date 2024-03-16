package handler

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForCart(firebaseApp validation.IFirebaseApp, cartRequests cart.IRequests, userRequests users.IRequests) {
	CartRouter := h.Router.Group("/api/cart")
	CartRouter.Use(firebaseMiddleware(firebaseApp))
	{
		CartRouter.Use(userMiddleware(userRequests))
		{
			CartRouter.GET("", func(ctx *gin.Context) {
				userId := ctx.GetString("userId")
				Cart, _ := cartRequests.Get(userId)
				ctx.JSON(http.StatusOK, gin.H{"items": Cart})
			})
			CartRouter.POST("", func(ctx *gin.Context) {
				userId := ctx.GetString("userId")

				cartRequestPayload, err := utils.GetPayloadFromBody(ctx, &cart.CartRequestPayload{})
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				log.Print(cartRequestPayload)
				err = cartRequests.Register(userId, cartRequestPayload)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully registered")
			})
			CartRouter.DELETE("", func(ctx *gin.Context) {
				userId := ctx.GetString("userId")
				itemId, err := utils.GetQuery("itemId", ctx)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				err = cartRequests.Delete(userId, *itemId)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, "Item was successfully deleted")
			})
		}
	}
}
