package manufacturer

import "github.com/gin-gonic/gin"

type ExampleManufacturerRequests struct {
}

func (m ExampleManufacturerRequests) RegisterItem(i ItemRegisterPayload, ctx *gin.Context) error {
	return nil
}

func (m ExampleManufacturerRequests) UpdateItem(i ItemUpdatePayload, ctx *gin.Context) error {
	return nil
}

func (m ExampleManufacturerRequests) DeleteItem(itemId string, ctx *gin.Context) error {
	return nil
}
