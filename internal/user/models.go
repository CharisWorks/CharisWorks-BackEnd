package user

import (
	"time"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserId       string       `json:"user_id" gorm:"user_id"`
	UserProfile  *UserProfile `json:"profile" gorm:"profile"`
	UserAddress  *UserAddress `json:"address" gorm:"address"`
	Manufacturer Manufacturer `json:"manufacturer" gorm:"manufacturer"`
}
type Manufacturer struct {
	StripeAccountId *string             `json:"stripe_account_id" gorm:"stripe_account_id"`
	Items           []items.ItemPreview `json:"items" gorm:"items"`
}
type UserProfile struct {
	DisplayName string    `json:"display_name" gorm:"display_name"`
	Description string    `json:"description" gorm:"description"`
	CreatedAt   time.Time `json:"crated_at" gorm:"created_at"`
}
type UserAddress struct {
	FirstName     string  `json:"first_name" gorm:"first_name"`
	FirstNameKana string  `json:"first_name_kana" gorm:"first_name_kana"`
	LastName      string  `json:"last_name" gorm:"last_name"`
	LastNameKana  string  `json:"last_name_kana" gorm:"last_name_kana"`
	ZipCode       string  `json:"zip_code" gorm:"zip_code"`
	Address1      string  `json:"address_1" gorm:"address_1"`
	Address2      string  `json:"address_2" gorm:"address_2"`
	Address3      *string `json:"address_3" gorm:"address_3"`
	PhoneNumber   string  `json:"phone_number" gorm:"phone_number"`
}
type UserProfileRegisterPayload struct {
	DisplayName string `json:"display_name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
type UserAddressRegisterPayload struct {
	FirstName     string  `json:"first_name" binding:"required"`
	FirstNameKana string  `json:"first_name_kana" binding:"required"`
	LastName      string  `json:"last_name" binding:"required"`
	LastNameKana  string  `json:"last_name_kana" binding:"required"`
	ZipCode       string  `json:"zip_code" binding:"required"`
	Address1      string  `json:"address_1" binding:"required"`
	Address2      string  `json:"address_2" binding:"required"`
	Address3      *string `json:"address_3"`
	PhoneNumber   string  `json:"phone_number" binding:"required"`
}

type HistoryUser struct {
	HistoryUserId string    `json:"history_user_id"`
	UserId        string    `json:"user_id"`
	DisplayName   string    `json:"display_name"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"crated_at"`
}

type IUserRequests interface {
	UserCreate(ctx *gin.Context, UserDB IUserDB) error
	UserGet(ctx *gin.Context, UserDB IUserDB) (*User, error)
	UserDelete(ctx *gin.Context, UserDB IUserDB) error
	UserProfileUpdate(*gin.Context, IUserDB) error
	UserAddressRegister(*gin.Context, IUserDB) error
	UserAddressUpdate(*gin.Context, IUserDB) error
}
type IUserUtils interface {
	InspectAddressRegisterPayload(UserAddressRegisterPayload) error
	InspectProfileUpdatePayload(UserProfile) error
	InspectAddressUpdatePayload(UserAddress) error
}
type IUserDB interface {
	CreateUser(UserId string, HistoryUserId int) error
	GetUser(UserId string) (*User, error)
	DeleteUser(UserId string) error
	UpdateProfile(string, UserProfile) error
	RegisterAddress(string, UserAddressRegisterPayload) error
	UpdateAddress(string, UserAddress) error
}
type IUserDBHistory interface {
	GetUser(UserId string) (*User, error)
	RegisterUserProfile(UserProfile UserProfile, UserId string) error
}
