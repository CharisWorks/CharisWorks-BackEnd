package users

type Requests struct {
	UserRepository IRepository
	UserUtils      IUtils
}

func (r Requests) UserCreate(userId string) error {
	err := r.UserRepository.CreateUser(userId)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) UserGet(userId string) (*User, error) {
	User, err := r.UserRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (r Requests) UserDelete(userId string) error {
	err := r.UserRepository.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) UserProfileUpdate(userId string, userProfile UserProfile) error {
	updatePayload := r.UserUtils.InspectProfileUpdatePayload(userProfile)
	err := r.UserRepository.UpdateProfile(userId, updatePayload)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) UserAddressRegister(userId string, userAddressRegisterPayload UserAddressRegisterPayload) error {
	payload, err := r.UserUtils.InspectAddressRegisterPayload(userAddressRegisterPayload)
	if err != nil {
		return err
	}
	err = r.UserRepository.RegisterAddress(userId, payload)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) UserAddressUpdate(userId string, userAddress UserAddress) error {
	updatePayload, err := r.UserUtils.InspectAddressUpdatePayload(userAddress)
	if err != nil {
		return err
	}
	err = r.UserRepository.UpdateAddress(userId, updatePayload)
	if err != nil {
		return err
	}
	return nil
}
