package transaction

import (
	"time"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/gin-gonic/gin"
)

type TransactionPreview struct {
	TransactionId int               `json:"transaction_id"`
	Items         []TransactionItem `json:"items"`
	TransactionAt time.Time         `json:"transaction_at"`
	Status        TransactionStatus `json:"status"`
}
type TransactionDetails struct {
	TransactionId int                `json:"transaction_id"`
	TrackingId    string             `json:"tracking_id"`
	UserAddress   TransactionAddress `json:"address"`
	Items         []TransactionItem  `json:"items"`
	TransactionAt time.Time          `json:"transaction_at"`
	Status        TransactionStatus  `json:"status"`
}
type TransactionStatus string

const (
	Pending  TransactionStatus = "Pending"
	Complete TransactionStatus = "Complete"
	Cancel   TransactionStatus = "Cancel"
	Fail     TransactionStatus = "Fail"
	Refund   TransactionStatus = "Refund"
)

type TransactionAddress struct {
	ZipCode     string `json:"zip_code"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	RealName    string `json:"real_name"`
}
type TransactionItem struct {
	ItemId   string `json:"item_id"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
}
type ITransactionRequests interface {
	GetList(userId string) (*map[int]TransactionPreview, error)
	GetDetails(userId string, transactionId string) (*TransactionDetails, error)
	Create(ctx *gin.Context, CartRequests cart.IRequests, cartRepository cart.IRepository, CartUtils cart.IUtils) error
}

type ITransactionStripeUtils interface {
	PurchaseComplete(StipeTransactionId string) error
	PurchaseCancel(StipeTransactionId string) error
	PurchaseFail(StipeTransactionId string) error
	PurchaseRefund(StipeTransactionId string) error
}

type ITransactionHistoryRepository interface {
	GetList(UserId string) (*[]TransactionPreview, error)
	GetDetails(TransactionId string) (*TransactionDetails, string, error)
	Register(UserId string, transactionDetails TransactionDetails) error
	StatusUpdate(string, TransactionStatus) error
}
