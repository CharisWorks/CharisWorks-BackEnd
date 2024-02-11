package user

import "github.com/gin-gonic/gin"

func UserGet(UserId string, i IUserRequests, ctx *gin.Context) (*User, error) {
	return i.UserGet(UserId, ctx), nil
}
