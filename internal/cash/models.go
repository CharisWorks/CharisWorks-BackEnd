package cash

import (
	"time"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/gin-gonic/gin"
)

type IStripeRequests interface {
	GetClientSecret(userId string, cartRequests cart.IRequests, cartRepository cart.IRepository, cartUtils cart.IUtils) (url *string, err error)
	GetRegisterLink(email string, user users.User, userRepository users.IRepository) (url *string, err error)
	GetStripeMypageLink(stripeAccountId string) (url *string, err error)
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
	GetTransactionList(ctx *gin.Context, transactionHistoryRepository ITransactionHistoryRepository) (*[]TransactionPreview, error)
	GetTransactionDetails(ctx *gin.Context) (*TransactionDetails, error)
	CreateTransaction(ctx *gin.Context, cartRequests cart.IRequests, cartRepository cart.IRepository, cartUtils cart.IUtils) error
}

type ITransactionStripeUtils interface {
	PurchaseComplete(stipeTransactionId string) error
	PurchaseCancel(stipeTransactionId string) error
	PurchaseFail(stipeTransactionId string) error
	PurchaseRefund(stipeTransactionId string) error
}

type ITransactionDB interface {
	ReduceStock(itemId string, quantity int) error
}

type ITransactionHistoryRepository interface {
	GetTransactionList(userId string) (*[]TransactionPreview, error)
	GetTransactionDetails(transactionId string) (*TransactionDetails, error)
	RegisterTransaction(userId string, transactionDetails TransactionDetails) (*string, error)
	TransactionStatusUpdate(string, TransactionStatus) error
}
