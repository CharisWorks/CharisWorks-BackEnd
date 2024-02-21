package cart

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

type Cart struct {
	ItemId         string                      `json:"item_id"`
	Quantity       int                         `json:"quantity"`
	ItemProperties items.ItemPreviewProperties `json:"properties"`
}
type CartRequestPayload struct {
	ItemId   string `json:"item_id" binding:"required" `
	Quantity int    `json:"quantity" binding:"required"`
}
type internalCart struct {
	Cart      Cart
	itemStock int
	status    string
}
type ICartRequests interface {
	Get(*gin.Context, ICartDB, string) (*[]Cart, error)
	Register(CartRequestPayload, ICartDB, *gin.Context) error
	Delete(string, *gin.Context) error
}
type ICartDB interface {
	GetCart(userId string) (*[]internalCart, error)
	RegisterCart(userId string, c CartRequestPayload) error
	UpdateCart(userId string, c CartRequestPayload) error
	DeleteCart(userId string, itemId string) error
}
type ICartUtils interface {
	InspectCart([]internalCart) (*[]internalCart, error)
	ConvertCart([]internalCart) *[]Cart
}
