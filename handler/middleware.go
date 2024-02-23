package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/user"
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
		ctx.Set("UserId", UserID)

		ctx.Next()
	}
}
func userMiddleware(UserRequests user.IUserRequests) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		EmailVerified := ctx.MustGet("EmailVerified").(bool)
		if !EmailVerified {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "email is not verified"})
			ctx.Abort()
			return
		}
		UserId := ctx.MustGet("UserId").(string)
		User, err := UserRequests.UserGet(UserId, ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			ctx.Abort()
			return
		}
		if User == nil {
			err := UserRequests.UserCreate(UserId, ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "create user for DB"})
			ctx.Abort()
			return
		}
		ctx.Set("User", User)
		//内部の実行タイミング
		ctx.Next()
	}
}
func stripeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		User := ctx.MustGet("User").(*user.User)
		if !User.UserProfile.IsManufacturer {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Account is not manufacturer"})
			ctx.Abort()
			return
		}
		//内部の実行タイミング
		ctx.Next()

	}
}

func manufacturerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		User := ctx.MustGet("User").(*user.User)
		if !User.UserProfile.IsManufacturer {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Account is not manufacturer"})
			ctx.Abort()
			return
		}
		if *User.Manufacturer.StripeAccountId == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "cannot get stripe account id"})
			ctx.Abort()
			return
		}
		ctx.Set("Stripe_Account_Id", User.Manufacturer.StripeAccountId)
		Account, err := cash.GetAccount(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "stripeのアカウントが取得できませんでした。"})
			ctx.Abort()
			return
		}
		if !Account.PayoutsEnabled {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "口座が登録されていません。"})
			ctx.Abort()
			return
		}
		ctx.Set("User", User)
		//内部の実行タイミング
		ctx.Next()

	}
}