package manufacturer

import "github.com/gin-gonic/gin"

type ItemRegisterPayload struct {
	Name    *string                     `json:"name" binding:"required"`
	Price   *int                        `json:"price" binding:"required"`
	Details *ItemRegisterDetailsPayload `json:"details" binding:"required"`
}
type ItemRegisterDetailsPayload struct {
	Status      *string   `json:"status" binding:"required"`
	Stock       *int      `json:"stock" binding:"required"`
	Size        *int      `json:"size" binding:"required"`
	Description *string   `json:"description" binding:"required"`
	Tags        *[]string `json:"tags" binding:"required"`
}

type ItemUpdatePayload struct {
	ItemId                      *string                     `json:"item_id" binding:"required"`
	ItemUpdatePropertiesPayload ItemUpdatePropertiesPayload `json:"properties" binding:"required"`
}
type ItemUpdatePropertiesPayload struct {
	Name    *string                   `json:"name"`
	Price   *int                      `json:"price"`
	Details *ItemUpdateDetailsPayload `json:"details"`
}

type ItemUpdateDetailsPayload struct {
	Status      *string   `json:"status"`
	Stock       *int      `json:"stock"`
	Size        *int      `json:"size"`
	Description *string   `json:"description" `
	Tags        *[]string `json:"tags"`
}

type IManufacturerRequests interface {
	RegisterItem(ItemRegisterPayload, *gin.Context) error
	UpdateItem(ItemUpdatePayload, *gin.Context) error
	DeleteItem(string, *gin.Context) error
}

type IManufactuerUtils interface {
	InspectRegisterPayload(ItemRegisterPayload) error
	InspectUpdatePayload(ItemUpdatePayload) error
}

type IManufacturerDB interface {
	RegisterItem(ItemRegisterPayload) error
	UpdateItem(ItemUpdatePayload) error
	DeleteItem(string) error
}
