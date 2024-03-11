package users

type Requests struct {
	UserRepository IRepository
	UserUtils      IUtils
}

func (r Requests) Create(userId string) error {
	err := r.UserRepository.Create(userId)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) Get(userId string) (*User, error) {
	User, err := r.UserRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (r Requests) Delete(userId string) error {
	err := r.UserRepository.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) ProfileUpdate(userId string, userProfile UserProfile) error {
	updatePayload := r.UserUtils.InspectProfileUpdatePayload(userProfile)
	err := r.UserRepository.UpdateProfile(userId, updatePayload)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) AddressRegister(userId string, AddressRegisterPayload AddressRegisterPayload) error {
	payload, err := r.UserUtils.InspectAddressRegisterPayload(AddressRegisterPayload)
	if err != nil {
		return err
	}
	err = r.UserRepository.RegisterAddress(userId, payload)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) AddressUpdate(userId string, userAddress UserAddress) error {
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
