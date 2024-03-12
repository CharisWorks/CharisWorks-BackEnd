package handler

import (
	"log"
	"net/http"

	stripe1 "github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"

	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func (h *Handler) SetupRoutesForStripe(firebaseApp validation.IFirebaseApp, UserRequests users.Requests, stripeRequests stripe1.IRequests, transactionRequests transaction.IRequests) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthGP4F3QjdR0SKk77E4pGHrsBAQEHia6lasXyujFOKXDyrodAxaE6PH6u2kNCVSdC5dBIRh82u00XqHQIZjM"
	StripeRouter := h.Router.Group("/api")
	StripeRouter.Use(firebaseMiddleware(firebaseApp))
	{
		StripeRouter.GET("/buy", func(ctx *gin.Context) {
			// レスポンスの処理
			userId := ctx.GetString("userId")
			ClientSecret, _, err := transactionRequests.Purchase(userId)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"clientSecret": ClientSecret})

		})
		StripeRouter.GET("/transaction", func(ctx *gin.Context) {
			userId := ctx.GetString("userId")
			TransactionList, err := transactionRequests.GetList(userId)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(http.StatusOK, TransactionList)
		})
		StripeRouter.GET("/transaction/:transactionId", func(ctx *gin.Context) {
			transactionId, err := utils.GetParams("transactionId", ctx)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			userId := ctx.GetString("userId")
			TransactionDetails, err := transactionRequests.GetDetails(userId, *transactionId)
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
		StripeManufacturerRouter.Use(userMiddleware(UserRequests))
		{
			StripeManufacturerRouter.GET("/create", func(ctx *gin.Context) {
				email := ctx.GetString("email")
				user, exist := ctx.Get(string(user))
				if !exist {
					err := utils.InternalError{Message: utils.InternalErrorUnAuthorized}
					ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
					return
				}
				URL, err := stripeRequests.GetRegisterLink(email, *user.(*users.User))
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"url": URL})

			})
			StripeManufacturerRouter.Use(stripeMiddleware())
			{
				StripeManufacturerRouter.GET("/mypage", func(ctx *gin.Context) {
					user, exist := ctx.Get(string(user))
					log.Print(user)
					if !exist {
						err := utils.InternalError{Message: utils.InternalErrorUnAuthorized}
						ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
						return
					}
					URL, err := stripeRequests.GetStripeMypageLink(user.(users.User).UserProfile.StripeAccountId)
					if err != nil {
						utils.ReturnErrorResponse(ctx, err)
						return
					}
					ctx.JSON(http.StatusOK, gin.H{"url": URL})
				})
			}
		}
	}

}
