package handler

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func (h *Handler) SetupRoutesForStripe(firebaseApp validation.IFirebaseApp) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthx8UK5v9CdaucJOZj3FwkOHZ8KjDt25IAvplosSab4uybQOyE2Ne6xxxI4Rnh8pWEbYUwPoPG00wvseAHzl"
	StripeRouter := h.Router.Group("/api")
	StripeRouter.Use(firebaseMiddleware(firebaseApp))

	{
		StripeRouter.GET("/buy", func(ctx *gin.Context) {
			// レスポンスの処理
			err := cash.CreatePaymentIntent(ctx, cash.ExampleTransactionUtils{}, cart.ExapleCartRequest{})
			if err != nil {
				return
			}

		})
		StripeRouter.GET("/transaction", func(ctx *gin.Context) {
			err := cash.GetTransactionList(ctx, cash.ExampleTransactionRequests{})
			if err != nil {
				return
			}
		})
		StripeRouter.GET("/transaction/:transactionId", func(ctx *gin.Context) {
			TransactionId, err := getQuery("transactionId", ctx)
			if err != nil {
				return
			}
			err = cash.GetTransactionDetails(*TransactionId, ctx, cash.ExampleTransactionRequests{})
			if err != nil {
				return
			}
		})
	}
	StripeManufacturerRouter := h.Router.Group("/api/stripe")
	StripeManufacturerRouter.Use(firebaseMiddleware(firebaseApp))
	{

		StripeManufacturerRouter.GET("/create", func(ctx *gin.Context) {
			cash.CreateStripeAccount(ctx)
		})
		StripeManufacturerRouter.GET("/mypage", func(ctx *gin.Context) {
			cash.GetMypage(ctx)
		})
	}

}
