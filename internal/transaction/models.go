package transaction

import (
	"time"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type TransactionPreview struct {
	TransactionId string            `json:"transaction_id"`
	Items         []TransactionItem `json:"items"`
	TransactionAt time.Time         `json:"transaction_at"`
	Status        TransactionStatus `json:"status"`
}
type TransactionDetails struct {
	TransactionId string             `json:"transaction_id"`
	TrackingId    string             `json:"tracking_id"`
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
}
type IRequests interface {
	GetList(userId string) (*[]TransactionPreview, error)
	GetDetails(userId string, transactionId string) (*TransactionDetails, error)
	Purchase(userId string) (*string, error)
	PurchaseRefund(stripeTransactionId string) error
}
type IWebhook interface {
	PurchaseComplete(stripeTransactionId string) error
	PurchaseFail(stripeTransactionId string) error
	PurchaseCanceled(stripeTransactionId string) error
}

type ITransactionRepository interface {
	GetList(UserId string) (*map[string]TransactionPreview, error)
	GetDetails(TransactionId string) (*TransactionDetails, string, error)
	Register(userId string, stripeTransactionId string, transactionId string, internalCartList []cart.InternalCart) error
	StatusUpdate(string, TransactionStatus) error
}
