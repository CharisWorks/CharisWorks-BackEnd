package user

type UserRequests struct {
}

func (u UserRequests) UserGet(UserId string) *User {
	user := ExampleUser(UserId)
	return &user
}
func (u UserRequests) UserDelete(UserId string) (message string) {
	return ""
}
func (u UserRequests) UserProfileRegister(UserProfile UserProfile) (message string) {
	return ""
}
func (u UserRequests) UserProfileUpdate(UserProfile UserProfile) (message string) {
	return ""
}
func (u UserRequests) UserAddressRegister(UserAddress UserAddress) (message string) {
	return ""
}
func (u UserRequests) UserAddressUpdate(UserAddress UserAddress) (message string) {
	return ""
}
func UserGet(UserId string, i IUserRequests) *User {
	return i.UserGet(UserId)
}
