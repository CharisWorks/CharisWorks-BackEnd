package cart

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
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
	Get(*gin.Context, ICartDB, ICartUtils, string) (*[]Cart, error)
	Register(CartRequestPayload, ICartDB, ICartUtils, *gin.Context, string) error
	Delete(string, ICartDB, ICartUtils, *gin.Context, string) error
}
type ICartDB interface {
	GetCart(userId string) (*[]InternalCart, error)
	RegisterCart(userId string, c CartRequestPayload) error
	UpdateCart(userId string, c CartRequestPayload) error
	DeleteCart(userId string, itemId string) error
	GetItem(itemId string) (*itemStatus, error)
}
type ICartUtils interface {
	InspectCart([]InternalCart) (map[string]InternalCart, error)
	ConvertCart(map[string]InternalCart) []Cart
	GetTotalAmount(map[string]InternalCart) int
	InspectPayload(c CartRequestPayload, itemStatus itemStatus) (*CartRequestPayload, error)
}
