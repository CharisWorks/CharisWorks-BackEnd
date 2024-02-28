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

// firstorinitをそのうち使うかもしれない
func (r UserDB) CreateUser(UserId string, historyUserId int) error {
	DBUser := new(utils.User)
	log.Print("UserId: ", UserId)
	log.Print("historyUserId: ", historyUserId)
	DBUser.Id = UserId
	DBUser.HistoryUserId = historyUserId
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
	if err := r.DB.Table("users").Where("id = ?", UserId).First(DBUser).Error; err != nil {
		return nil, err
	}
	user := new(User)
	user.UserId = DBUser.Id
	user.UserProfile = UserProfile{
		DisplayName: DBUser.DisplayName,
		Description: DBUser.Description,
		CreatedAt:   DBUser.CreatedAt,
	}
	log.Print("successfully got data. id: ", DBUser.Id, "displayname: ", DBUser.DisplayName, "description: ", DBUser.Description, "created_at: ", DBUser.CreatedAt)
	user.Manufacturer = Manufacturer{
		StripeAccountId: &DBUser.StripeAccountId,
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
		Address3:      &Address.Address_3,
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
	if err := r.DB.Table("users").Where("id = ?", UserId).Update("id", "deleted_"+UserId+"_"+time.Now().String()).Error; err != nil {
		return err
	}
	return nil
}

func (r UserDB) UpdateProfile(UserId string, payload map[string]interface{}) error {
	for key, value := range payload {
		log.Print("key: ", key, "value: ", value)
	}
	log.Print("UserId: ", UserId)
	if err := r.DB.Table("users").Where("id = ?", UserId).Updates(payload).Error; err != nil {
		return err
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
	if err := r.DB.Create(Shipping).Error; err != nil {
		return err
	}

	return nil
}
func (r UserDB) UpdateAddress(UserId string, payload map[string]string) error {
	if err := r.DB.Table("shippings").Where("id = ?", UserId).Updates(payload).Error; err != nil {
		return err
	}
	return nil
}
