package users

import (
	"time"
)

type User struct {
	UserId      string      `json:"user_id" gorm:"user_id"`
	UserProfile UserProfile `json:"profile" gorm:"profile"`
	UserAddress UserAddress `json:"address" gorm:"address"`
}
type UserProfile struct {
	DisplayName     string    `json:"display_name" gorm:"display_name"`
	Description     string    `json:"description" gorm:"description"`
	StripeAccountId string    `json:"stripe_account_id" gorm:"stripe_account_id"`
	CreatedAt       time.Time `json:"crated_at" gorm:"created_at"`
}
type UserAddress struct {
	FirstName     string `json:"first_name" gorm:"first_name"`
	FirstNameKana string `json:"first_name_kana" gorm:"first_name_kana"`
	LastName      string `json:"last_name" gorm:"last_name"`
	LastNameKana  string `json:"last_name_kana" gorm:"last_name_kana"`
	ZipCode       string `json:"zip_code" gorm:"zip_code"`
	Address1      string `json:"address_1" gorm:"address_1"`
	Address2      string `json:"address_2" gorm:"address_2"`
	Address3      string `json:"address_3" gorm:"address_3"`
	PhoneNumber   string `json:"phone_number" gorm:"phone_number"`
}
type UserProfileRegisterPayload struct {
	DisplayName string `json:"display_name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
type AddressRegisterPayload struct {
	FirstName     string `json:"first_name" binding:"required"`
	FirstNameKana string `json:"first_name_kana" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	LastNameKana  string `json:"last_name_kana" binding:"required"`
	ZipCode       string `json:"zip_code" binding:"required"`
	Address1      string `json:"address_1" binding:"required"`
	Address2      string `json:"address_2" binding:"required"`
	Address3      string `json:"address_3"`
	PhoneNumber   string `json:"phone_number" binding:"required"`
}

type IRequests interface {
	Create(userId string) error
	Get(userId string) (*User, error)
	Delete(userId string) error
	ProfileUpdate(userId string, userProfile UserProfile) error
	AddressRegister(userId string, AddressRegisterPayload AddressRegisterPayload) error
	AddressUpdate(userId string, userAddress UserAddress) error
}
type IUtils interface {
	InspectAddressRegisterPayload(AddressRegisterPayload) (AddressRegisterPayload, error)
	InspectProfileUpdatePayload(UserProfile) map[string]interface{}
	InspectAddressUpdatePayload(UserAddress) (map[string]interface{}, error)
}
type IRepository interface {
	Create(UserId string) error
	Get(UserId string) (user *User, err error)
	Delete(UserId string) error
	UpdateProfile(string, map[string]interface{}) error
	RegisterAddress(string, AddressRegisterPayload) error
	UpdateAddress(string, map[string]interface{}) error
}
