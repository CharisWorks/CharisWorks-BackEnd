package manufacturer

import (
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type ItemRegisterPayload struct {
	Name    string                     `json:"name" binding:"required"`
	Price   int                        `json:"price" binding:"required"`
	Details ItemRegisterDetailsPayload `json:"details" binding:"required"`
}
type ItemRegisterDetailsPayload struct {
	Stock       int      `json:"stock" binding:"required"`
	Size        int      `json:"size" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Tags        []string `json:"tags" binding:"required" `
}

type ItemUpdatePayload struct {
	ItemId                      string                      `json:"item_id" binding:"required"`
	ItemUpdatePropertiesPayload ItemUpdatePropertiesPayload `json:"properties" binding:"required"`
}
type ItemUpdatePropertiesPayload struct {
	Name    string                   `json:"name"`
	Price   int                      `json:"price"`
	Details ItemUpdateDetailsPayload `json:"details"`
}

type ItemUpdateDetailsPayload struct {
	Status      string   `json:"status"`
	Stock       int      `json:"stock"`
	Size        int      `json:"size"`
	Description string   `json:"description" `
	Tags        []string `json:"tags"`
}

type IItemRequests interface {
	Register(itemRegisterPayload ItemRegisterPayload, userId string) error
	Update(query map[string]interface{}) error
	Delete(itemId string) error
}

type IInspectPayloadUtils interface {
	Register(ItemRegisterPayload) error
	Update(map[string]interface{}) (*map[string]interface{}, error)
}

type IItemRepository interface {
	Register(itemId string, i ItemRegisterPayload, userId string) error
	Update(i map[string]interface{}, itemId string) error
	Delete(itemId string) error
}
type IHistoryRepository interface {
	Register(i utils.Item) (err error)
	Get(itemId string) (utils.Item, error)
}
type IHistoryUtils interface {
	HistoryItemUpdate(i utils.Item, payload map[string]interface{}) (utils.Item, error)
}
