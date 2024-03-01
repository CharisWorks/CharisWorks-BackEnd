package utils

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetPayloadFromBody[T any](ctx *gin.Context, p *T) (*T, error) {
	bind := new(T)
	err := ctx.BindJSON(&bind)
	if err != nil {
		return nil, &InternalError{Message: InternalErrorInvalidPayload}
	}
	return bind, nil
}

func GetQuery(params string, isRequired bool, ctx *gin.Context) (*string, error) {
	itemId := ctx.Query(params)
	if itemId == "" {
		return nil, &InternalError{Message: InternalErrorInvalidQuery}
	}
	return &itemId, nil
}

func GetParams(params string, isRequired bool, ctx *gin.Context) (*string, error) {
	itemId := ctx.Param(params)
	if itemId == "" {
		return nil, &InternalError{Message: InternalErrorInvalidQuery}
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
