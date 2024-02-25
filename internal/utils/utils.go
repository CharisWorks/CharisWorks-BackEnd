package utils

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetPayloadFromBody[T any](ctx *gin.Context, p *T) (*T, error) {
	bind := new(T)
	err := ctx.BindJSON(&bind)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "The request payload is malformed or contains invalid data."})
		return nil, &InternalError{Message: InternalErrorInvalidPayload}
	}
	return bind, nil
}

func GetQuery(params string, isRequired bool, ctx *gin.Context) (*string, error) {
	itemId := ctx.Query(params)
	if itemId == "" {
		if isRequired {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot get" + params})
			return nil, &InternalError{Message: InternalErrorInvalidQuery}
		}
		return nil, &InternalError{Message: InternalErrorInvalidQuery}
	}
	return &itemId, nil
}

func GetParams(params string, isRequired bool, ctx *gin.Context) (*string, error) {
	itemId := ctx.Param(params)
	if itemId == "" {
		if isRequired {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot get" + params})
			return nil, &InternalError{Message: InternalErrorInvalidParams}
		}
		return nil, nil
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
