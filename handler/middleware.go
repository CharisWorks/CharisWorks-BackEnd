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
			log.Print(err)
			utils.AbortContextWithError(ctx, err)
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
			err := &utils.InternalError{Message: utils.InternalErrorInvalidUserRequest}
			utils.AbortContextWithError(ctx, err)
			return
		}
		if !EmailVerified.(bool) {
			err := &utils.InternalError{Message: utils.InternalErrorEmailIsNotVerified}
			utils.AbortContextWithError(ctx, err)
			return
		}
		userId := ctx.GetString(string(userId))
		User, err := UserRequests.Get(userId)
		if err != nil {
			if err.Error() == string(utils.InternalErrorNotFound) && User.UserId == "" {
				log.Print("creating user for DB")
				err := UserRequests.Create(userId)
				if err != nil {
					utils.AbortContextWithError(ctx, err)
					return
				}
				err = &utils.InternalError{Message: utils.InternalErrorNotFound}
				utils.AbortContextWithError(ctx, err)
				return
			}
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
			return
		}
		if User.(*users.User).UserProfile.StripeAccountId == "" {
			err := utils.InternalError{Message: utils.InternalErrorAccountIsNotSatisfied}
			ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
			return
		}
		//内部の実行タイミング
		ctx.Next()

	}
}

func manufacturerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		User := ctx.MustGet(string(user)).(users.User)
		if User.UserProfile.StripeAccountId == "" {
			err := &utils.InternalError{Message: utils.InternalErrorAccountIsNotSatisfied}
			utils.AbortContextWithError(ctx, err)
			return
		}
		Account, err := cash.GetAccount(User.UserProfile.StripeAccountId)
		if err != nil {
			utils.AbortContextWithError(ctx, err)
			return
		}
		if !Account.PayoutsEnabled {
			err := &utils.InternalError{Message: utils.InternalErrorManufacturerDoesNotHaveBank}
			utils.AbortContextWithError(ctx, err)
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
