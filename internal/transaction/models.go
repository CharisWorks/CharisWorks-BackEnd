package transaction

import (
	"time"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type TransactionPreview struct {
	TransactionId string            `json:"transaction_id"`
	Email         string            `json:"email"`
	Items         []TransactionItem `json:"items"`
	TransactionAt time.Time         `json:"transaction_at"`
	Status        TransactionStatus `json:"status"`
}
type TransactionDetails struct {
	TransactionId string             `json:"transaction_id"`
	TrackingId    string             `json:"tracking_id"`
	Email         string             `json:"email"`
	TotalAmount   int                `json:"total_amount"`
	TotalPrice    int                `json:"total_price"`
	UserAddress   TransactionAddress `json:"address"`
	Items         []TransactionItem  `json:"items"`
	TransactionAt time.Time          `json:"transaction_at"`
	Status        TransactionStatus  `json:"status"`
}
type InternalTransactionDetails struct {
	TransactionId string                    `gorm:"transaction_id"`
	UserAddress   utils.Shipping            `gorm:"embedded"`
	Items         []InternalTransactionItem `gorm:"embedded"`
}

type TransactionStatus string

const (
	Pending   TransactionStatus = "Pending"
	Complete  TransactionStatus = "Complete"
	Cancelled TransactionStatus = "Cancel"
	Fail      TransactionStatus = "Fail"
	Refund    TransactionStatus = "Refund"
)

type TransactionAddress struct {
	ZipCode     string `json:"zip_code"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	RealName    string `json:"real_name"`
}
type TransactionItem struct {
	ItemId     string `json:"item_id"`
	Quantity   int    `json:"quantity"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	TransferId string `json:"transfer_id"`
	Status     string `json:"status"`
}
type InternalTransactionItem struct {
	ItemId                  string `gorm:"item_id"`
	Price                   int    `gorm:"price"`
	Name                    string `gorm:"name"`
	Quantity                int    `gorm:"quantity"`
	Description             string `gorm:"description"`
	Tags                    string `gorm:"tags"`
	ManufacturerUserId      string `gorm:"manufacturer_user_id"`
	ManufacturerName        string `gorm:"manufacturer_name"`
	ManufacturerDescription string `gorm:"manufacturer_description"`
	TransferId              string `gorm:"transfer_id"`
	Status                  string `gorm:"status"`
}
type IRequests interface {
	GetList(userId string) ([]TransactionPreview, error)
	GetDetails(userId string, transactionId string) (TransactionDetails, error)
	Purchase(userId string, email string) (clientSecret string, transactionId string, err error)
	PurchaseRefund(stripeTransferId string, transactionId string) error
}
type IWebhook interface {
	PurchaseComplete(stripeTransactionId string) (t TransactionDetails, err error)
	PurchaseFail(stripeTransactionId string) error
	PurchaseCanceled(stripeTransactionId string) error
}
type transfer struct {
	amount          int
	itemId          string
	stripeAccountId string
	transferId      string
}

// test complete
type IRepository interface {
	GetList(UserId string) (map[string]TransactionPreview, error)
	GetDetails(transactionId string) (TransactionDetails, string, []transfer, error)
	Register(userId string, email string, transactionId string, internalCartList []cart.InternalCart) error
	StatusUpdate(transactionId string, conditions map[string]interface{}) error
	StatusUpdateItems(transactionId string, itemId string, conditions map[string]interface{}) error
}
