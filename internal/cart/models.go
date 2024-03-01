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
	Status CartItemStatus `json:"status"`
}
type CartItemStatus string

const (
	CartItemStatusAvailable   CartItemStatus = "Available"
	CartItemStatusStockOver   CartItemStatus = "Stock over"
	CartItemStatusNoStock     CartItemStatus = "No stock"
	CartItemStatusInvalidItem CartItemStatus = "Invalid item"
)

type CartRequestPayload struct {
	ItemId   string `json:"item_id" binding:"required" `
	Quantity int    `json:"quantity" binding:"required"`
}
type InternalCart struct {
	Index     int
	Cart      Cart
	ItemStock int
	Status    items.ItemStatus
}
type itemStatus struct {
	itemStock int
	status    items.ItemStatus
}
type ICartRequests interface {
	Get(userId string, CartDB ICartDB, CartUtils ICartUtils) (*[]Cart, error)
	Register(userId string, cartRequestPayload CartRequestPayload, CartDB ICartDB, CartUtils ICartUtils) error
	Delete(userId string, itemId string, CartDB ICartDB, CartUtils ICartUtils) error
}
type ICartDB interface {
	GetCart(UserId string) (*[]InternalCart, error)
	RegisterCart(UserId string, c CartRequestPayload) error
	UpdateCart(UserId string, c CartRequestPayload) error
	DeleteCart(UserId string, itemId string) error
	GetItem(itemId string) (*itemStatus, error)
}
type ICartUtils interface {
	InspectCart([]InternalCart) (map[string]InternalCart, error)
	ConvertCart(map[string]InternalCart) []Cart
	GetTotalAmount(map[string]InternalCart) int
	InspectPayload(c CartRequestPayload, itemStatus itemStatus) (*CartRequestPayload, error)
}
