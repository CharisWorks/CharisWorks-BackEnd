package cash

import (
	"time"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/user"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/gin-gonic/gin"
)

type IStripeRequests interface {
	GetClientSecret(*gin.Context, cart.ICartRequests, cart.ICartDB, cart.ICartUtils, string) (url *string, err error)
	GetRegisterLink(*gin.Context) (url *string, err error)
	GetStripeMypageLink(*gin.Context) (url *string, err error)
}

type TransactionPreview struct {
	TransactionId string             `json:"transaction_id"`
	Items         []TransactionItems `json:"items"`
	TransactionAt time.Time          `json:"transaction_at"`
}
type TransactionDetails struct {
	TransactionId string             `json:"transaction_id"`
	TrackingId    string             `json:"tracking_id"`
	UserAddress   TransactionAddress `json:"address"`
	Items         []TransactionItems `json:"items"`
	TransactionAt time.Time          `json:"transaction_at"`
}
type TransactionAddress struct {
	ZipCode     string `json:"zip_code"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	RealName    string `json:"real_name"`
}
type TransactionItems struct {
	TransactionId string `json:"transaction_id"`
	ItemId        string `json:"item_id"`
	Quantity      int    `json:"quantity"`
}
type ITransactionRequests interface {
	GetTransactionList(*gin.Context, user.IUserDB, items.IItemDB) ([]TransactionPreview, error)
	GetTransactionDetails(ctx *gin.Context, TransactionId string) (TransactionDetails, error)
}
type ITransactionDB interface {
	PurchaseComplete(TransactionDetails) (string, error)
}
type ITransactionDBhistory interface {
	GetItem(ItemId string) (cart.Cart, error)
	GetTransactionList(UserId string) ([]TransactionPreview, error)
	GetTransactionDetails(TransactionId string) (TransactionDetails, error)
	CreateTransaction(UserId string, transactionDetails TransactionDetails) (string, error)
}
