package cart

import "github.com/charisworks/charisworks-backend/internal/items"

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
	Get() (*[]Cart, error)
	Register(c CartRequestPayload) error
	Update(c CartRequestPayload) error
	Delete(ItemId string) error
}
