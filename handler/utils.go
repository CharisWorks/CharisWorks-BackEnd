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

func firebaseMiddleware(app validation.FirebaseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		UID, err := app.VerifyIDToken(c, c.Request.Header.Get("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
		}
		c.Set("UserId", UID)
		//内部の実行タイミング
		c.Next()

	}
}
func manufacturerMiddleware(app validation.FirebaseApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		User, err := user.UserGet(c.MustGet("UserId").(string), user.ExampleUserRequests{})
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
		}
		if User == nil || User.Manufacturer == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Account is not manufacturer"})
			c.Abort()
		}

		//内部の実行タイミング
		c.Next()

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

func getPayloadFromBody[T any](c *gin.Context, p *T) (*T, error) {
	bind := new(T)
	err := c.ShouldBindJSON(&bind)
	if err != nil {
		err := new(internalError)
		err.badRequestErrorForPayload("The request payload is malformed or contains invalid data.")
		return nil, err
	}
	return bind, nil
}
