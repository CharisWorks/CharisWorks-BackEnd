package user

import (
	"log"
	"regexp"
	"time"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (r UserDB) CreateUser(UserId string, historyUserId int) error {
	DBUser := new(utils.User)
	log.Print("UserId: ", UserId)
	log.Print("historyUserId: ", historyUserId)
	DBUser.Id = UserId
	DBUser.CreatedAt = time.Now()
	result := r.DB.Create(DBUser)
	log.Print("result: ", result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r UserDB) GetUser(UserId string) (*User, error) {
	DBUser := new(utils.User)
	result := r.DB.Table("users").Where("id = ?", UserId).First(DBUser)
	if result.Error != nil {
		return nil, result.Error
	}
	user := new(User)
	user.UserId = DBUser.Id
	user.UserProfile = &UserProfile{
		DisplayName: DBUser.DisplayName,
		Description: DBUser.Description,
		CreatedAt:   DBUser.CreatedAt,
	}
	user.Manufacturer = Manufacturer{
		StripeAccountId: &DBUser.StripeAccountId,
	}
	Address := new(utils.Shipping)
	_ = r.DB.Table("shippings").Where("id = ?", UserId).First(Address)
	user.UserAddress = &UserAddress{
		FirstName:     Address.FirstName,
		FirstNameKana: Address.FirstNameKana,
		LastName:      Address.LastName,
		LastNameKana:  Address.LastNameKana,
		ZipCode:       Address.ZipCode,
		Address1:      Address.Address1,
		Address2:      Address.Address2,
		Address3:      &Address.Address3,
		PhoneNumber:   Address.PhoneNumber,
	}
	log.Print("user: ", user)
	regex := regexp.MustCompile(`acct_\w+`)
	matches := regex.FindAllString(*user.Manufacturer.StripeAccountId, -1)
	for _, match := range matches {
		if regex.MatchString(match) {
			items := new([]items.ItemPreview)
			_ = r.DB.Table("items").Where("manufacturer_user_id = ?", UserId).Find(items)
			user.Manufacturer.Items = *items
		}
	}
	return user, nil
}
func (r UserDB) DeleteUser(UserId string) error {
	result := r.DB.Table("users").Where("id = ?", UserId).Update("id", "deleted_"+UserId+"_"+time.Now().String())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r UserDB) UpdateProfile(UserId string, payload UserProfile) error {
	result := r.DB.Table("users").Where("id = ?", UserId).Update("display_name", payload.DisplayName).Update("description", payload.Description)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r UserDB) RegisterAddress(UserId string, payload UserAddressRegisterPayload) error {
	Shipping := new(utils.Shipping)
	Shipping.Id = UserId
	Shipping.ZipCode = payload.ZipCode
	Shipping.Address1 = payload.Address1
	Shipping.Address2 = payload.Address2
	if payload.Address3 != nil {
		Shipping.Address3 = *payload.Address3
	}
	Shipping.PhoneNumber = payload.PhoneNumber
	result := r.DB.Create(Shipping)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r UserDB) UpdateAddress(UserId string, payload UserAddress) error {
	result := r.DB.Table("shippings").Where("id = ?", UserId).Update("zip_code", payload.ZipCode).Update("address1", payload.Address1).Update("address2", payload.Address2).Update("address3", payload.Address3).Update("phone_number", payload.PhoneNumber)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
