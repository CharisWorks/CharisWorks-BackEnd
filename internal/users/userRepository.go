package users

import (
	"log"
	"time"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

// firstorinitをそのうち使うかもしれない
func (r UserDB) CreateUser(UserId string, historyUserId int) error {
	DBUser := new(utils.User)
	DBUser.Id = UserId
	DBUser.HistoryUserId = historyUserId
	DBUser.CreatedAt = time.Now()
	result := r.DB.Create(DBUser)
	if err := result.Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (r UserDB) GetUser(UserId string) (*User, error) {
	DBUser := new(utils.User)
	if err := r.DB.Table("users").Where("id = ?", UserId).First(DBUser).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	user := new(User)
	user.UserId = DBUser.Id
	user.UserProfile = UserProfile{
		DisplayName:     DBUser.DisplayName,
		Description:     DBUser.Description,
		StripeAccountId: DBUser.StripeAccountId,
		CreatedAt:       DBUser.CreatedAt,
	}
	Address := new(utils.Shipping)
	_ = r.DB.Table("shippings").Where("id = ?", UserId).First(Address)
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
func (r UserDB) DeleteUser(UserId string) error {
	if err := r.DB.Table("users").Where("id = ?", UserId).Delete(utils.User{}).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}

func (r UserDB) UpdateProfile(UserId string, payload map[string]interface{}) error {
	if err := r.DB.Table("users").Where("id = ?", UserId).Updates(payload).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (r UserDB) RegisterAddress(UserId string, payload UserAddressRegisterPayload) error {
	Shipping := new(utils.Shipping)
	Shipping.Id = UserId
	Shipping.ZipCode = payload.ZipCode
	Shipping.Address_1 = payload.Address1
	Shipping.Address_2 = payload.Address2
	if payload.Address3 != nil {
		Shipping.Address_3 = *payload.Address3
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
func (r UserDB) UpdateAddress(UserId string, payload map[string]interface{}) error {
	if err := r.DB.Table("shippings").Where("id = ?", UserId).Updates(payload).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
