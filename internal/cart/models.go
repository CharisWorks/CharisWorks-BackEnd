package cart

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

type Cart struct {
	ItemId                string                      `json:"item_id"`
	Quantity              int                         `json:"quantity"`
	ItemPreviewProperties items.ItemPreviewProperties `json:"properties"`
}
type CartRequestPayload struct {
	ItemId   string `json:"item_id" binding:"required" `
	Quantity int    `json:"quantity" binding:"required"`
}
type ICartRequest interface {
	Get(*gin.Context) (*[]Cart, error)
	Register(CartRequestPayload, *gin.Context) error
	Update(CartRequestPayload, *gin.Context) error
	Delete(string, *gin.Context) error
}
