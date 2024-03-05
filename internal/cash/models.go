package cash

import (
	"time"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/gin-gonic/gin"
)

type IStripeRequests interface {
	GetClientSecret(userId string, CartRequests cart.ICartRequests, cartRepository cart.ICartRepository, CartUtils cart.ICartUtils) (url *string, err error)
	GetRegisterLink(email string, user users.User, UserDB users.IUserRepository) (url *string, err error)
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
type TransactionItems struct {
	TransactionId string `json:"transaction_id"`
	ItemId        string `json:"item_id"`
	Quantity      int    `json:"quantity"`
}
type ITransactionRequests interface {
	GetList(ctx *gin.Context, TransactionDBHistory ITransactionDBHistory) (*[]TransactionPreview, error)
	GetDetails(ctx *gin.Context) (*TransactionDetails, error)
	Create(ctx *gin.Context, CartRequests cart.ICartRequests, cartRepository cart.ICartRepository, CartUtils cart.ICartUtils) error
}

type ITransactionStripeUtils interface {
	PurchaseComplete(StipeTransactionId string) error
	PurchaseCancel(StipeTransactionId string) error
	PurchaseFail(StipeTransactionId string) error
	PurchaseRefund(StipeTransactionId string) error
}

type ITransactionDB interface {
	ReduceStock(itemId string, Quantity int) error
	ChangeState(itemId string, State items.Status)
}

type ITransactionDBHistory interface {
	GetList(UserId string) (*[]TransactionPreview, error)
	GetDetails(TransactionId string) (*TransactionDetails, error)
	Register(UserId string, transactionDetails TransactionDetails) (*string, error)
	StatusUpdate(string, TransactionStatus) error
}
