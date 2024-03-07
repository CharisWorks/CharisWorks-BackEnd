package cash

import (
	"github.com/charisworks/charisworks-backend/internal/users"
)

type IStripeRequests interface {
	CreatePaymentintent(userId string, totalAmount int) (ClientSecret *string, StripeTransactionId *string, err error)
	GetRegisterLink(email string, user users.User) (url *string, err error)
	GetStripeMypageLink(stripeAccountId string) (url *string, err error)
}
