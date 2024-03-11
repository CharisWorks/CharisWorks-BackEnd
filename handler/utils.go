package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	Router *gin.Engine
}

func NewHandler(router *gin.Engine) *Handler {
	return &Handler{
		Router: router,
	}
}

type structInContext string

const (
	userId        structInContext = "userId"
	user          structInContext = "user"
	emailVerified structInContext = "EmailVerified"
	stripeEvent   structInContext = "Event"
)
