package handler

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
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
	err := ctx.ShouldBindJSON(&bind)
	if err != nil {
		err := new(internalError)
		err.badRequestErrorForPayload("The request payload is malformed or contains invalid data.")
		ctx.JSON(http.StatusBadRequest, err)
		return nil, err
	}
	return bind, nil
}
