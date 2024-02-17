package cash

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type IStripeRequests interface {
	GetClientSecret(amount decimal.Decimal, email string) (url string, err error)
	GetRegisterLink(email string) (url string, err error)
	GetStripeMypageLink(email string) (url string, err error)
}
type PurchasePayload struct {
	Amount int    `json:"amount" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

type TransactionPreview struct {
	TransactionId string      `json:"transaction_id"`
	Items         []cart.Cart `json:"items"`
}
type TransactionDetails struct {
	TransactionId string           `json:"transaction_id"`
	TrackingId    string           `json:"tracking_id"`
	UserAddress   user.UserAddress `json:"address"`
	Items         []cart.Cart      `json:"items"`
}
type ITransactionRequests interface {
	GetTransactionList(*gin.Context) ([]TransactionPreview, error)
	GetTransactionDetails(ctx *gin.Context, TransactionId string) (TransactionDetails, error)
}
type ITransactionUtils interface {
	GetTotalAmount([]cart.Cart) int64
	InspectCart([]cart.Cart) error
}
