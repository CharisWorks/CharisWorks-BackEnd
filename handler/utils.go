package handler

import (
	"net/http"
	"time"

	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router *gin.Engine
}

func NewHandler(router *gin.Engine) *Handler {
	return &Handler{
		Router: router,
	}
}

func firebaseMiddleware(app validation.IFirebaseApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UserID, err := app.VerifyIDToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}
		ctx.Set("UserId", UserID)
		User, err := user.UserGet(UserID, user.ExampleUserRequests{}, ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}
		ctx.Set("User", *User)
		//内部の実行タイミング
		ctx.Next()

	}
}
func manufacturerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//ctx.Set("Stripe_Account_Id", "acct_1OkZRtPMQkfESzTI")
		//ctx.Set("Stripe_Account_Id", "acct_1Okj9YPFjznovTf3")
		User := ctx.MustGet("User").(user.User)
		if !User.UserProfile.IsManufacturer {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Account is not manufacturer"})
			ctx.Abort()
			return
		}
		if User.Manufacturer.StripeAccountId == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "cannot get stripe account id"})
			ctx.Abort()
			return
		}
		ctx.Set("Stripe_Account_Id", User.Manufacturer.StripeAccountId)
		Account, err := cash.GetAcount(ctx)
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

type internalError struct {
	Message string
}

func (e *internalError) Error() string {
	return e.Message
}
func (e *internalError) setError(msg string) {
	e.Message = msg
}

func getPayloadFromBody[T any](ctx *gin.Context, p *T) (*T, error) {
	bind := new(T)
	err := ctx.BindJSON(&bind)
	if err != nil {
		err := new(internalError)
		err.setError("The request payload is malformed or contains invalid data.")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}
	return bind, nil
}

func getQuery(params string, ctx *gin.Context) (*string, error) {
	itemId := ctx.Query(params)
	if itemId == "" {
		err := new(internalError)
		err.setError("cannot get" + params)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}
	return &itemId, nil
}

func getParams(params string, ctx *gin.Context) (*string, error) {
	itemId := ctx.Param(params)
	if itemId == "" {
		err := new(internalError)
		err.setError("cannot get" + params)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}
	return &itemId, nil
}
func CORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PATCH",
			"DELETE",
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Content-Type",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Authorization",
			"Access-Control-Allow-Credentials",
		},

		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

}
