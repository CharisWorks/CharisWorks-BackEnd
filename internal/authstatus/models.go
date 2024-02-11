package authstatus

import "github.com/gin-gonic/gin"

type IAuthStatusRequests interface {
	Check(string, *gin.Context) (bool, error)
}
type Email struct {
	Email string `json:"email" binding:"required"`
}
