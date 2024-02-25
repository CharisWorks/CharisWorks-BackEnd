package user

import "github.com/gin-gonic/gin"

type UserRequests struct {
}

func (r UserRequests) UserCreate(UserId string, ctx *gin.Context, UserDB IUserDB) error {
	UserDB.CreateUser(UserId, 1)
	return nil
}
func (r UserRequests) UserGet(UserId string, ctx *gin.Context, UserDB IUserDB) (*User, error) {
	User, err := UserDB.GetUser(UserId)
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (r UserRequests) UserDelete(UserId string, ctx *gin.Context) error {
	return nil
}
func (r UserRequests) UserProfileRegister(p UserProfileRegisterPayload, ctx *gin.Context) error {
	return nil
}
func (r UserRequests) UserProfileUpdate(p UserProfile, ctx *gin.Context) error {
	return nil
}
func (r UserRequests) UserAddressRegister(p UserAddressRegisterPayload, ctx *gin.Context) error {
	return nil
}
func (r UserRequests) UserAddressUpdate(p UserAddress, ctx *gin.Context) error {
	return nil
}
