package handler

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
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
		User, err := UserRequests.UserGet(userId)
		if err != nil {
			utils.ReturnErrorResponse(ctx, err)
			ctx.Abort()
			return
		}
		if User == nil {
			log.Print("creating user for DB")
			err := UserRequests.UserCreate(userId)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				ctx.Abort()
				return
			}
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
			ctx.JSON(utils.Code(utils.InternalMessage(err.Error())), gin.H{"message": err.Error()})
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
		User, exist := ctx.Get("User")
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
