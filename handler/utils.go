package handler

import (
	"net/http"
	"time"

	"github.com/charisworks/charisworks-backend/internal/utils"
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

func getPayloadFromBody[T any](ctx *gin.Context, p *T) (*T, error) {
	bind := new(T)
	err := ctx.BindJSON(&bind)
	if err != nil {
		err := new(utils.InternalError)
		err.SetError("The request payload is malformed or contains invalid data.")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}
	return bind, nil
}

func getQuery(params string, ctx *gin.Context) (*string, error) {
	itemId := ctx.Query(params)
	if itemId == "" {
		err := new(utils.InternalError)
		err.SetError("cannot get" + params)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}
	return &itemId, nil
}

func getParams(params string, ctx *gin.Context) (*string, error) {
	itemId := ctx.Param(params)
	if itemId == "" {
		err := new(utils.InternalError)
		err.SetError("cannot get" + params)
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
