package cart

import (
	"github.com/charisworks/charisworks-backend/internal/items"
)

type Cart struct {
	ItemId         string                    `json:"item_id"`
	Quantity       int                       `json:"quantity"`
	ItemProperties CartItemPreviewProperties `json:"properties"`
}
type CartItemPreviewProperties struct {
	Name    string                 `json:"name"`
	Price   int                    `json:"price"`
	Details CartItemPreviewDetails `json:"details"`
}

type CartItemPreviewDetails struct {
	Status ItemStatus `json:"status"`
}
type ItemStatus string

const (
	Available   ItemStatus = "Available"
	StockOver   ItemStatus = "Stock over"
	NoStock     ItemStatus = "No stock"
	InvalidItem ItemStatus = "Invalid item"
)

type CartRequestPayload struct {
	ItemId   string `json:"item_id" binding:"required" `
	Quantity int    `json:"quantity" binding:"required"`
}
type InternalCart struct {
	Index     int
	Cart      Cart
	Item      InternalItem
	ItemStock int
	Status    items.Status
}
type InternalItem struct {
	Price                   int      `gorm:"price"`
	Name                    string   `gorm:"name"`
	Description             string   `gorm:"description"`
	Tags                    []string `gorm:"tags"`
	Size                    int      `gorm:"size"`
	ManufacturerUserId      string   `gorm:"manufacturer_user_id"`
	ManufacturerName        string   `gorm:"manufacturer_name"`
	ManufacturerDescription string   `gorm:"manufacturer_description"`
	ManufacturerStripeId    string   `gorm:"manufacturer_stripe_account_id"`
}
type IRequests interface {
	Get(userId string) ([]Cart, error)
	Register(userId string, cartRequestPayload CartRequestPayload) error
	Delete(userId string, itemId string) error
}
type IRepository interface {
	Get(UserId string) ([]InternalCart, error)
	Register(UserId string, c CartRequestPayload) error
	Update(UserId string, c CartRequestPayload) error
	Delete(UserId string, itemId string) error
	DeleteAll(UserId string) error
}
type IUtils interface {
	Inspect([]InternalCart) (map[string]InternalCart, error)
	Convert(map[string]InternalCart) []Cart
	GetTotalAmount(map[string]InternalCart) int
	InspectPayload(c CartRequestPayload, itemStatus items.ItemStatus) (*CartRequestPayload, error)
}
