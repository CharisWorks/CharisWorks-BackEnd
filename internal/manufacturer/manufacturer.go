package manufacturer

import "github.com/gin-gonic/gin"

func RegisterItem(p ItemRegisterPayload, i IManufacturerRequests, ctx *gin.Context) error {
	err := i.RegisterItem(p, ctx)
	return err
}
func UpdateItem(p ItemUpdatePayload, i IManufacturerRequests, ctx *gin.Context) error {
	err := i.UpdateItem(p, ctx)
	return err
}
func DeleteItem(itemId string, i IManufacturerRequests, ctx *gin.Context) error {
	err := i.DeleteItem(itemId, ctx)
	return err
}
