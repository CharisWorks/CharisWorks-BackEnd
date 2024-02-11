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
	StripeAccountId string              `json:"stripe_account_id"`
	Items           []items.ItemPreview `json:"items"`
}
type UserProfile struct {
	DisplayName string    `json:"display_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"crated_at"`
}
type UserAddress struct {
	RealName    string  `json:"real_name"`
	ZipCode     string  `json:"zip_code"`
	Address1    string  `json:"address_1"`
	Address2    string  `json:"address_2"`
	Address3    *string `json:"address_3"`
	PhoneNumber string  `json:"phone_number"`
}
type UserProfileRegisterPayload struct {
	DisplayName string `json:"display_name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
type UserAddressRegisterPayload struct {
	RealName    string `json:"real_name" binding:"required"`
	ZipCode     string `json:"zip_code" binding:"required"`
	Address1    string `json:"address_1" binding:"required"`
	Address2    string `json:"address_2" binding:"required"`
	Address3    string `json:"address_3"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
type IUserRequests interface {
	UserGet(string, *gin.Context) *User
	UserDelete(string, *gin.Context) error
	UserProfileRegister(UserProfileRegisterPayload, *gin.Context) error
	UserProfileUpdate(UserProfile, *gin.Context) error
	UserAddressRegister(UserAddressRegisterPayload, *gin.Context) error
	UserAddressUpdate(UserAddress, *gin.Context) error
}
