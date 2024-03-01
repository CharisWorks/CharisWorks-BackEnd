package manufacturer

import "github.com/gin-gonic/gin"

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

type IManufacturerRequests interface {
	RegisterItem(*gin.Context) error
	UpdateItem(*gin.Context) error
	DeleteItem(*gin.Context) error
}

type IManufactuerUtils interface {
	InspectRegisterPayload(ItemRegisterPayload) error
	InspectUpdatePayload(map[string]interface{}) (*map[string]interface{}, error)
}

type IManufacturerDB interface {
	RegisterItem(i ItemRegisterPayload, history_item_id string, userId string) error
	UpdateItem(i map[string]interface{}, history_item_id string) error
	DeleteItem(itemId string) error
}
type IManufactuerDBHistory interface {
	RegisterItemHistory(i ItemRegisterPayload) (history_item_id string, err error)
}
