package stripe

import "github.com/shopspring/decimal"

type IStripeRequests interface {
	GetPurchaseLink(amount decimal.Decimal, email string) (url string, err error)
	GetRegisterLink(email string) (url string, err error)
	GetStripeMypageLink(email string) (url string, err error)
}
type PurchasePayload struct {
	Amount int    `json:"amount" binding:"required"`
	Email  string `json:"email" binding:"required"`
}
