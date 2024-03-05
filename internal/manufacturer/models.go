package manufacturer

import (
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
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

type IRequests interface {
	RegisterItem(itemRegisterPayload ItemRegisterPayload, historyItemId int, userId string, manufacturerDB IRepository, manufacturerUtils IUtils) error
	UpdateItem(map[string]interface{}, IRepository, IUtils) error
	DeleteItem(*gin.Context) error
}

type IUtils interface {
	InspectRegisterPayload(ItemRegisterPayload) error
	InspectUpdatePayload(map[string]interface{}) (*map[string]interface{}, error)
}

type IRepository interface {
	RegisterItem(i ItemRegisterPayload, userId string) error
	UpdateItem(i map[string]interface{}, history_item_id int) error
	DeleteItem(itemId string) error
}
type IHistoryRepository interface {
	HistoryItemRegister(i utils.Item) (history_item_id int, err error)
	HistoryItemGet(itemId string) (utils.Item, error)
}
type IHistoryUtils interface {
	HistoryItemUpdate(i utils.Item, payload map[string]interface{}) (utils.Item, error)
}
