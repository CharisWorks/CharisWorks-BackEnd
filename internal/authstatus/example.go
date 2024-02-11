package authstatus

import "github.com/gin-gonic/gin"

func ExampleAuthStatus(email string) bool {
	return true
}

type ExampleAuthStatusRequests struct {
}

func (a ExampleAuthStatusRequests) Check(email string, ctx *gin.Context) (bool, error) {
	return ExampleAuthStatus(email), nil
}
