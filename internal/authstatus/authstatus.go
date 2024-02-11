package authstatus

import "github.com/gin-gonic/gin"

func AuthStatusCheck(email Email, i IAuthStatusRequests, ctx *gin.Context) (bool, error) {
	return i.Check(email.Email, ctx)
}
