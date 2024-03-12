package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76/webhook"
)

func firebaseMiddleware(app validation.IFirebaseApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UserID, err := app.VerifyIDToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}
		ctx.Set(string(userId), UserID)

		ctx.Next()
	}
}
func userMiddleware(UserRequests users.IRequests) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		EmailVerified, exist := ctx.Get(string(emailVerified))
		if !exist {
			err := utils.InternalError{Message: utils.InternalErrorInvalidUserRequest}
			ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		if !EmailVerified.(bool) {
			err := utils.InternalError{Message: utils.InternalErrorEmailIsNotVerified}
			ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		userId := ctx.GetString(string(userId))
		User, err := UserRequests.Get(userId)
		if err != nil {
			if err.Error() == string(utils.InternalErrorNotFound) {
				log.Print("creating user for DB")
				err := UserRequests.Create(userId)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					ctx.Abort()
					return
				}
				err = &utils.InternalError{Message: utils.InternalErrorNotFound}
				ctx.JSON(http.StatusOK, gin.H{"message": "new user"})
				ctx.Abort()
				return
			}
			utils.ReturnErrorResponse(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set(string(user), User)
		//内部の実行タイミング
		ctx.Next()
	}
}
func stripeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		User, exist := ctx.Get(string(user))
		if !exist {
			err := utils.InternalError{Message: utils.InternalErrorInvalidUserRequest}
			ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		if User.(*users.User).UserProfile.StripeAccountId == "" {
			err := utils.InternalError{Message: utils.InternalErrorAccountIsNotSatisfied}
			ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		//内部の実行タイミング
		ctx.Next()

	}
}

func manufacturerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		User := ctx.MustGet("User").(*users.User)
		if User.UserProfile.StripeAccountId == "" {
			err := utils.InternalError{Message: utils.InternalErrorAccountIsNotSatisfied}
			ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		Account, err := cash.GetAccount(User.UserProfile.StripeAccountId)
		if err != nil {
			utils.ReturnErrorResponse(ctx, err)
			ctx.Abort()
			return
		}
		if !Account.PayoutsEnabled {
			err := utils.InternalError{Message: utils.InternalErrorManufacturerDoesNotHaveBank}
			ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set(string(user), User)
		//内部の実行タイミング
		ctx.Next()

	}
}

func webhookMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const MaxBodyBytes = int64(65536)
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, MaxBodyBytes)

		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			ctx.Abort()
			return
		}

		// Pass the request body and Stripe-Signature header to ConstructEvent, along with the webhook signing key
		// You can find your endpoint's secret in your webhook settings
		endpointSecret := os.Getenv("STRIPE_KEY")
		log.Print(string(body), "ctx:", ctx.Request.Header)
		event, err := webhook.ConstructEvent(body, ctx.Request.Header.Get("Stripe-Signature"), endpointSecret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
			ctx.Abort() // Return a 400 error on a bad signature
			return
		}
		ctx.Set(string(stripeEvent), event)
		log.Print("Webhook received!")
		ctx.Next()
	}
}
