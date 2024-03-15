package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
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

func GetQuery(params string, ctx *gin.Context) (*string, error) {
	itemId := ctx.Query(params)
	if itemId == "" {
		return nil, &InternalError{Message: InternalErrorInvalidQuery}
	}
	return &itemId, nil
}

func GetParams(params string, ctx *gin.Context) (*string, error) {
	itemId := ctx.Param(params)
	if itemId == "" {
		return nil, &InternalError{Message: InternalErrorInvalidParams}
	}
	return &itemId, nil
}
func CORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"http://localhost:3001",
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
func GenerateRandomString() string {
	// Snowflakeノード設定
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	// Snowflake ID生成
	id := node.Generate().Int64()

	// 使用する文字セット
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 文字列生成
	var sb strings.Builder
	base := int64(len(charSet))
	for id > 0 {
		sb.WriteByte(charSet[id%base])
		id /= base
	}

	// 文字列を反転させる（SnowflakeのIDは末尾がよりランダムとなるため）
	str := sb.String()
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ConvertToJST(utcTime time.Time) string {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// タイムゾーンの読み込みに失敗した場合のエラーハンドリング
		fmt.Println("タイムゾーンの読み込みに失敗しました:", err)
		return ""
	}
	jstTime := utcTime.In(loc)
	return jstTime.Format("2006-01-02")
}
