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

func (h *Handler) SetupRoutesForStripe(firebaseApp validation.IFirebaseApp, transactionRequests cash.ITransactionRequests, StripeRequests cash.IStripeRequests) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthx8UK5v9CdaucJOZj3FwkOHZ8KjDt25IAvplosSab4uybQOyE2Ne6xxxI4Rnh8pWEbYUwPoPG00wvseAHzl"
	StripeRouter := h.Router.Group("/api")
	StripeRouter.Use(firebaseMiddleware(firebaseApp))
	{
		StripeRouter.GET("/buy", func(ctx *gin.Context) {
			// レスポンスの処理
			ClientSecret, err := StripeRequests.GetClientSecret(ctx, cart.ExampleCartDB{}, cart.CartUtils{})
			if err != nil {
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"clientSecret": ClientSecret})

		})
		StripeRouter.GET("/transaction", func(ctx *gin.Context) {
			TransactionList, err := transactionRequests.GetTransactionList(ctx)
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
			TransactionDetails, err := transactionRequests.GetTransactionDetails(ctx, *TransactionId)
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
				URL, err := StripeRequests.GetRegisterLink(ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"url": URL})

			})
			StripeManufacturerRouter.GET("/mypage", func(ctx *gin.Context) {
				URL, err := StripeRequests.GetStripeMypageLink(ctx)
				if err != nil {
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"url": URL})
			})
		}
	}

}
