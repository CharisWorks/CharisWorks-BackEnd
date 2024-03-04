package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func (h *Handler) SetupRoutesForStripe(firebaseApp validation.IFirebaseApp, transactionRequests cash.ITransactionRequests, StripeRequests cash.IStripeRequests, CartRequests cart.ICartRequests, cartDB cart.ICartRepository, cartUtils cart.ICartUtils, ItemDB items.IItemDB, TransactionDBHistory cash.ITransactionDBHistory, UserRequests users.IUserRequests, UserDB users.IUserDB) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthx8UK5v9CdaucJOZj3FwkOHZ8KjDt25IAvplosSab4uybQOyE2Ne6xxxI4Rnh8pWEbYUwPoPG00wvseAHzl"
	StripeRouter := h.Router.Group("/api")
	StripeRouter.Use(firebaseMiddleware(firebaseApp))
	{
		StripeRouter.GET("/buy", func(ctx *gin.Context) {
			// レスポンスの処理
			userId := ctx.GetString("userId")
			ClientSecret, err := StripeRequests.GetClientSecret(userId, CartRequests, cartDB, cartUtils)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"clientSecret": ClientSecret})

		})
		StripeRouter.GET("/transaction", func(ctx *gin.Context) {
			TransactionList, err := transactionRequests.GetList(ctx, TransactionDBHistory)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(http.StatusOK, TransactionList)
		})
		StripeRouter.GET("/transaction/:transactionId", func(ctx *gin.Context) {
			TransactionDetails, err := transactionRequests.GetDetails(ctx)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(http.StatusOK, TransactionDetails)
		})
	}
	StripeManufacturerRouter := h.Router.Group("/api/stripe")
	StripeManufacturerRouter.Use(firebaseMiddleware(firebaseApp))
	{
		StripeManufacturerRouter.Use(userMiddleware(UserRequests, UserDB))
		StripeManufacturerRouter.Use(stripeMiddleware())
		{
			StripeManufacturerRouter.GET("/create", func(ctx *gin.Context) {
				email := ctx.GetString("email")
				user, exist := ctx.Get("userId")
				if !exist {
					err := utils.InternalError{Message: utils.InternalErrorUnAuthorized}
					ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
					return
				}
				URL, err := StripeRequests.GetRegisterLink(email, user.(users.User), UserDB)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"url": URL})

			})
			StripeManufacturerRouter.GET("/mypage", func(ctx *gin.Context) {
				user, exist := ctx.Get("userId")
				if !exist {
					err := utils.InternalError{Message: utils.InternalErrorUnAuthorized}
					ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
					return
				}
				URL, err := StripeRequests.GetStripeMypageLink(user.(users.User).UserProfile.StripeAccountId)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"url": URL})
			})
		}
	}

}
