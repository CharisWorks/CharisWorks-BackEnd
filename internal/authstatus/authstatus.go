package authstatus

import "github.com/gin-gonic/gin"

func AuthStatusCheck(email Email, i IAuthStatusRequests, c *gin.Context) (bool, error) {
	return i.Check(email.Email, c)
}
