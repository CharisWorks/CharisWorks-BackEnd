package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func (h *Handler) SetupRoutesForStripe(firebaseApp validation.IFirebaseApp, transactionImpl cash.ITransactionRequests) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthx8UK5v9CdaucJOZj3FwkOHZ8KjDt25IAvplosSab4uybQOyE2Ne6xxxI4Rnh8pWEbYUwPoPG00wvseAHzl"
	StripeRouter := h.Router.Group("/api")
	StripeRouter.Use(firebaseMiddleware(firebaseApp))
	{
		StripeRouter.GET("/buy", func(ctx *gin.Context) {
			// レスポンスの処理
			ClientSecret, err := cash.CreatePaymentIntent(ctx, cash.ExampleTransactionUtils{}, cart.ExapleCartRequest{})
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"clientSecret": ClientSecret})

		})
		StripeRouter.GET("/transaction", func(ctx *gin.Context) {
			TransactionList, err := transactionImpl.GetTransactionList(ctx)
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, TransactionList)
		})
		StripeRouter.GET("/transaction/:transactionId", func(ctx *gin.Context) {
			TransactionId, err := getQuery("transactionId", ctx)
			if err != nil {
				return
			}
			TransactionDetails, err := transactionImpl.GetTransactionDetails(ctx, *TransactionId)
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, TransactionDetails)
		})
	}
	StripeManufacturerRouter := h.Router.Group("/api/stripe")
	StripeManufacturerRouter.Use(firebaseMiddleware(firebaseApp))
	{
		StripeManufacturerRouter.Use(userMiddleware(user.ExampleUserRequests{}))
		StripeManufacturerRouter.Use(stripeMiddleware())
		{
			StripeManufacturerRouter.GET("/create", func(ctx *gin.Context) {
				URL, err := cash.CreateStripeAccount(ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"url": URL})

			})
			StripeManufacturerRouter.GET("/mypage", func(ctx *gin.Context) {
				URL, err := cash.GetMypage(ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"url": URL})
			})
		}
	}

}
