package user

import (
	"log"
	"time"

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
	log.Print("user: ", user)
	return user, nil
}
func (r UserDB) DeleteUser(UserId string) error {
	result := r.DB.Table("users").Where("id = ?", UserId).Update("id", "deleted_"+UserId+"_"+time.Now().String())
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r UserDB) RegisterProfile(UserId string, payload UserProfileRegisterPayload) error {

	return nil
}

func (r UserDB) UpdateProfile(UserId string, payload UserProfile) error {

	return nil
}
func (r UserDB) RegisterAddress(UserId string, payload UserAddressRegisterPayload) error {

	return nil
}
func (r UserDB) UpdateAddress(UserId string, payload UserAddress) error {

	return nil
}
