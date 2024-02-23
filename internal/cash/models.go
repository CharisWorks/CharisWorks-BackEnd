package cash

import (
	"time"

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
	Status        TransactionStatus  `json:"status"`
}
type TransactionDetails struct {
	TransactionId string             `json:"transaction_id"`
	TrackingId    string             `json:"tracking_id"`
	UserAddress   TransactionAddress `json:"address"`
	Items         []TransactionItems `json:"items"`
	TransactionAt time.Time          `json:"transaction_at"`
	Status        TransactionStatus  `json:"status"`
}
type TransactionStatus string

const (
	TransactionStatusPending  TransactionStatus = "Pending"
	TransactionStatusComplete TransactionStatus = "Complete"
	TransactionStatusCancel   TransactionStatus = "Cancel"
	TransactionStatusFail     TransactionStatus = "Fail"
	TransactionStatusRefund   TransactionStatus = "Refund"
)

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
	GetTransactionList(ctx *gin.Context, TransactionDBHistory ITransactionDBHistory, userId string) (*[]TransactionPreview, error)
	GetTransactionDetails(ctx *gin.Context, TransactionId string) (*TransactionDetails, error)
	CreateTransaction(ctx *gin.Context, CartRequests cart.ICartRequests, CartDB cart.ICartDB, CartUtils cart.ICartUtils, userId string) error
}

type ITransactionStripeUtils interface {
	PurchaseComplete(StipeTransactionId string) error
	PurchaseCancel(StipeTransactionId string) error
	PurchaseFail(StipeTransactionId string) error
	PurchaseRefund(StipeTransactionId string) error
}

type ITransactionDB interface {
	ReduceStock(itemId string, Quantity int) error
}

type ITransactionDBHistory interface {
	GetTransactionList(UserId string) (*[]TransactionPreview, error)
	GetTransactionDetails(TransactionId string) (*TransactionDetails, error)
	RegisterTransaction(UserId string, transactionDetails TransactionDetails) (*string, error)
	TransactionStatusUpdate(string, TransactionStatus) error
}
