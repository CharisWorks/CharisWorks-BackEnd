package cart

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

type Cart struct {
	ItemId         int                       `json:"item_id"`
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
	ItemId   int `json:"item_id" binding:"required" `
	Quantity int `json:"quantity" binding:"required"`
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
	Get(*gin.Context, ICartDB, ICartUtils) (*[]Cart, error)
	Register(ICartDB, ICartUtils, *gin.Context) error
	Delete(ICartDB, ICartUtils, *gin.Context) error
}
type ICartDB interface {
	GetCart(UserId string) (*[]InternalCart, error)
	RegisterCart(UserId string, c CartRequestPayload) error
	UpdateCart(UserId string, c CartRequestPayload) error
	DeleteCart(UserId string, itemId int) error
	GetItem(itemId int) (*itemStatus, error)
}
type ICartUtils interface {
	InspectCart([]InternalCart) (map[int]InternalCart, error)
	ConvertCart(map[int]InternalCart) []Cart
	GetTotalAmount(map[int]InternalCart) int
	InspectPayload(c CartRequestPayload, itemStatus itemStatus) (*CartRequestPayload, error)
}
