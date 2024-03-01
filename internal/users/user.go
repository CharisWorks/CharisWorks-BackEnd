package users

type UserRequests struct {
}

func (r UserRequests) UserCreate(userId string, UserDB IUserDB) error {
	err := UserDB.CreateUser(userId, 1)
	if err != nil {
		return err
	}
	return nil
}
func (r UserRequests) UserGet(userId string, UserDB IUserDB) (*User, error) {
	User, _, err := UserDB.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (r UserRequests) UserDelete(userId string, UserDB IUserDB) error {
	err := UserDB.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
func (r UserRequests) UserProfileUpdate(userId string, userProfile UserProfile, UserDB IUserDB, UserUtils IUserUtils) error {
	updatePayload := UserUtils.InspectProfileUpdatePayload(userProfile)
	err := UserDB.UpdateProfile(userId, updatePayload)
	if err != nil {
		return err
	}
	return nil
}
func (r UserRequests) UserAddressRegister(userId string, userAddressRegisterPayload UserAddressRegisterPayload, userDB IUserDB, UserUtils IUserUtils) error {
	payload, err := UserUtils.InspectAddressRegisterPayload(userAddressRegisterPayload)
	if err != nil {
		return err
	}
	err = userDB.RegisterAddress(userId, payload)
	if err != nil {
		return err
	}
	return nil
}
func (r UserRequests) UserAddressUpdate(userId string, userAddress UserAddress, UserDB IUserDB, UserUtils IUserUtils) error {
	updatePayload, err := UserUtils.InspectAddressUpdatePayload(userAddress)
	if err != nil {
		return err
	}
	err = UserDB.UpdateAddress(userId, updatePayload)
	if err != nil {
		return err
	}
	return nil
}
