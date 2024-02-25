package manufacturer

import (
	"github.com/gin-gonic/gin"
)

type ExampleManufacturerRequests struct {
}

func (m ExampleManufacturerRequests) RegisterItem(ctx *gin.Context) error {
	return nil
}

func (m ExampleManufacturerRequests) UpdateItem(ctx *gin.Context) error {
	return nil
}

func (m ExampleManufacturerRequests) DeleteItem(ctx *gin.Context) error {
	return nil
}
