package handler

import (
	"log"
	"net/http"
	"time"

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
		UID, err := app.VerifyIDToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			ctx.Abort()
		}
		ctx.Set("UserId", UID)
		//内部の実行タイミング
		ctx.Next()

	}
}
func manufacturerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		User, err := user.UserGet(ctx.MustGet("UserId").(string), user.ExampleUserRequests{}, ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}
		if User == nil || User.Manufacturer == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Account is not manufacturer"})
			ctx.Abort()
			return
		}

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
func (e *internalError) badRequestErrorForPayload(msg string) {
	e.Message = msg
}

func getPayloadFromBody[T any](ctx *gin.Context, p *T) (*T, error) {
	bind := new(T)
	err := ctx.BindJSON(&bind)
	if err != nil {
		err := new(internalError)
		err.badRequestErrorForPayload("The request payload is malformed or contains invalid data.")
		ctx.JSON(http.StatusBadRequest, err)
		return nil, err
	}
	return bind, nil
}

func CORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"*",
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
	log.Print(r)
}
