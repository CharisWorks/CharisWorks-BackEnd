package user

import (
	"time"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserId       string        `json:"user_id"`
	UserProfile  *UserProfile  `json:"profile"`
	UserAddress  *UserAddress  `json:"address"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}
type Manufacturer struct {
	StripeAccountId *string             `json:"stripe_account_id"`
	Items           []items.ItemPreview `json:"items"`
}
type UserProfile struct {
	DisplayName    string `json:"display_name"`
	Description    string `json:"description"`
	IsManufacturer bool
	CreatedAt      time.Time `json:"crated_at"`
}
type UserAddress struct {
	FirstName     string  `json:"first_name"`
	FirstNameKana string  `json:"first_name_kana"`
	LastName      string  `json:"last_name"`
	LastNameKana  string  `json:"last_name_kana"`
	ZipCode       string  `json:"zip_code"`
	Address1      string  `json:"address_1"`
	Address2      string  `json:"address_2"`
	Address3      *string `json:"address_3"`
	PhoneNumber   string  `json:"phone_number"`
}
type UserProfileRegisterPayload struct {
	DisplayName string `json:"display_name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
type UserAddressRegisterPayload struct {
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

type HistoryUser struct {
	HistoryUserId string    `json:"history_user_id"`
	UserId        string    `json:"user_id"`
	DisplayName   string    `json:"display_name"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"crated_at"`
}

type IUserRequests interface {
	UserCreate(UserId string, ctx *gin.Context, UserDB IUserDB) error
	UserGet(UserID string, ctx *gin.Context, UserDB IUserDB) (*User, error)
	UserDelete(UserId string, ctx *gin.Context) error
	UserProfileRegister(UserProfileRegisterPayload, *gin.Context) error
	UserProfileUpdate(UserProfile, *gin.Context) error
	UserAddressRegister(UserAddressRegisterPayload, *gin.Context) error
	UserAddressUpdate(UserAddress, *gin.Context) error
}
type IUserUtils interface {
	InspectProfileRegisterPayload(UserProfileRegisterPayload) error
	InspectAddressRegisterPayload(UserAddressRegisterPayload) error
	InspectProfileUpdatePayload(UserProfile) error
	InspectAddressUpdatePayload(UserAddress) error
}
type IUserDB interface {
	CreateUser(UserId string, HistoryUserId int) error
	GetUser(UserId string) (*User, error)
	DeleteUser(UserId string) error
	RegisterProfile(string, UserProfileRegisterPayload) error
	UpdateProfile(string, UserProfile) error
	RegisterAddress(string, UserAddressRegisterPayload) error
	UpdateAddress(string, UserAddress) error
}
type IUserDBHistory interface {
	GetUser(UserId string) (*User, error)
	RegisterUserProfile(UserProfile UserProfile, UserId string) error
}
