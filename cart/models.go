package cart

import "github.com/charisworks/charisworks-backend/items"

type Cart struct {
	ItemId                string                      `json:"item_id"`
	ItemPreviewProperties items.ItemPreviewProperties `json:"properties"`
}
type CartRequestPayload struct {
	ItemId   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}
type ICartRequest interface {
	Get() *[]Cart
	Register(c CartRequestPayload) (message string)
	Update(c CartRequestPayload) (message string)
	Delete(ItemId string) (message string)
}
