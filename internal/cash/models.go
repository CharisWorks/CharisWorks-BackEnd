package cash

import (
	"github.com/charisworks/charisworks-backend/internal/users"
)

type IRequests interface {
	CreatePaymentintent(userId string, totalAmount int) (ClientSecret string, StripeTransactionId string, err error)
	GetRegisterLink(email string, user users.User) (url string, err error)
	GetStripeMypageLink(stripeAccountId string) (url string, err error)
}
type IUtils interface {
	Transfer(amount int, transactionId string, stripeAccountId string) (transferId *string)
	Refund(amount int, transactionId string, stripeAccountId string) (err error)
}
