package users

type UserRequests struct {
}

func (r UserRequests) UserCreate(userId string, UserDB IUserRepository) error {
	err := UserDB.CreateUser(userId)
	if err != nil {
		return err
	}
	return nil
}
func (r UserRequests) UserGet(userId string, UserDB IUserRepository) (*User, error) {
	User, err := UserDB.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (r UserRequests) UserDelete(userId string, UserDB IUserRepository) error {
	err := UserDB.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
func (r UserRequests) UserProfileUpdate(userId string, userProfile UserProfile, UserDB IUserRepository, UserUtils IUserUtils) error {
	updatePayload := UserUtils.InspectProfileUpdatePayload(userProfile)
	err := UserDB.UpdateProfile(userId, updatePayload)
	if err != nil {
		return err
	}
	return nil
}
func (r UserRequests) UserAddressRegister(userId string, userAddressRegisterPayload UserAddressRegisterPayload, userDB IUserRepository, UserUtils IUserUtils) error {
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
func (r UserRequests) UserAddressUpdate(userId string, userAddress UserAddress, UserDB IUserRepository, UserUtils IUserUtils) error {
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
