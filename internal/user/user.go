package user

func UserGet(UserId string, i IUserRequests) (*User, error) {
	return i.UserGet(UserId), nil
}
