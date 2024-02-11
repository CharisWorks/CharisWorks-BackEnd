package user

import "github.com/gin-gonic/gin"

func UserGet(UserId string, i IUserRequests, ctx *gin.Context) (*User, error) {
	return i.UserGet(UserId, ctx), nil
}

func UserDelete(UserId string, i IUserRequests, ctx *gin.Context) error {
	return i.UserDelete(UserId, ctx)
}

func UserProfileRegister(p UserProfileRegisterPayload, i IUserRequests, ctx *gin.Context) error {
	return i.UserProfileRegister(p, ctx)
}

func UserProfileUpdate(p UserProfile, i IUserRequests, ctx *gin.Context) error {
	return i.UserProfileUpdate(p, ctx)
}

func UserAddressRegister(p UserAddressRegisterPayload, i IUserRequests, ctx *gin.Context) error {
	return i.UserAddressRegister(p, ctx)
}

func UserAddressUpdate(p UserAddress, i IUserRequests, ctx *gin.Context) error {
	return i.UserAddressUpdate(p, ctx)
}
