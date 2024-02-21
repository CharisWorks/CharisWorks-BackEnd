package cash

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/gin-gonic/gin"
)

type IStripeRequests interface {
	GetClientSecret(*gin.Context, cart.ICartRequests, cart.ICartDB, cart.ICartUtils) (url *string, err error)
	GetRegisterLink(*gin.Context) (url *string, err error)
	GetStripeMypageLink(*gin.Context) (url *string, err error)
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
