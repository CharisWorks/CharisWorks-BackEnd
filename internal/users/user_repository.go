package users

import (
	"log"
	"time"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// firstorinitをそのうち使うかもしれない
func (r UserRepository) Create(UserId string) error {
	DBUser := new(utils.User)
	DBUser.Id = UserId
	DBUser.CreatedAt = time.Now()
	result := r.DB.Create(DBUser)
	if err := result.Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (r UserRepository) Get(UserId string) (user User, err error) {
	DBUser := new(utils.User)
	user = *new(User)
	if err := r.DB.Table("users").Where("id = ?", UserId).First(&DBUser).Error; err != nil {
		log.Print("DB error: ", err)
		if err.Error() == string(utils.InternalErrorNotFound) {
			return user, err
		}
		return user, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	user.UserId = DBUser.Id
	user.UserProfile = UserProfile{
		DisplayName:     DBUser.DisplayName,
		Description:     DBUser.Description,
		StripeAccountId: DBUser.StripeAccountId,
		CreatedAt:       DBUser.CreatedAt,
	}
	Address := new(utils.Shipping)
	_ = r.DB.Table("shippings").Where("id = ?", UserId).First(&Address)
	user.UserAddress = UserAddress{
		FirstName:     Address.FirstName,
		FirstNameKana: Address.FirstNameKana,
		LastName:      Address.LastName,
		LastNameKana:  Address.LastNameKana,
		ZipCode:       Address.ZipCode,
		Address1:      Address.Address_1,
		Address2:      Address.Address_2,
		Address3:      Address.Address_3,
		PhoneNumber:   Address.PhoneNumber,
	}
	return user, nil
}
func (r UserRepository) Delete(UserId string) error {
	if err := r.DB.Table("users").Where("id = ?", UserId).Delete(utils.User{}).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}

func (r UserRepository) UpdateProfile(UserId string, payload map[string]interface{}) error {
	if err := r.DB.Table("users").Where("id = ?", UserId).Updates(payload).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (r UserRepository) RegisterAddress(UserId string, payload AddressRegisterPayload) error {
	s := new(utils.Shipping)
	if err := r.DB.Table("shippings").Where("id = ?", UserId).First(&s).Error; err != nil {
		log.Print("DB error: ", err)
		if err.Error() != string(utils.InternalErrorNotFound) {
			return &utils.InternalError{Message: utils.InternalErrorDB}
		}
	} else {
		return &utils.InternalError{Message: utils.InternalErrorInvalidUserRequest}
	}
	Shipping := new(utils.Shipping)
	Shipping.Id = UserId
	Shipping.ZipCode = payload.ZipCode
	Shipping.Address_1 = payload.Address1
	Shipping.Address_2 = payload.Address2
	if payload.Address3 != "" {
		Shipping.Address_3 = payload.Address3
	}
	Shipping.PhoneNumber = payload.PhoneNumber
	Shipping.FirstName = payload.FirstName
	Shipping.FirstNameKana = payload.FirstNameKana
	Shipping.LastName = payload.LastName
	Shipping.LastNameKana = payload.LastNameKana
	if err := r.DB.Table("shippings").Create(Shipping).Error; err != nil {
		log.Print("DB error", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}

	return nil
}
func (r UserRepository) UpdateAddress(UserId string, payload map[string]interface{}) error {
	if err := r.DB.Table("shippings").Where("id = ?", UserId).Updates(payload).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
