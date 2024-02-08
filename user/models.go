package user

import (
	"time"

	"github.com/charisworks/charisworks-backend/items"
)

type User struct {
	UserId       string        `json:"user_id"`
	UserProfile  *UserProfile  `json:"profile"`
	UserAddress  *UserAddress  `json:"address"`
	Manufacturer *Manufacturer `json:"manufacturer"`
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

type Manufacturer struct {
	StripeAccountId string              `json:"stripe_account_id"`
	Items           []items.ItemPreview `json:"items"`
}
