package manufacturer

import "github.com/gin-gonic/gin"

func RegisterItem(p ItemRegisterPayload, i IManufacturerRequests, ctx *gin.Context) error {
	return i.RegisterItem(p, ctx)
}
func UpdateItem(p ItemUpdatePayload, i IManufacturerRequests, ctx *gin.Context) error {
	return i.UpdateItem(p, ctx)
}
func DeleteItem(itemId string, i IManufacturerRequests, ctx *gin.Context) error {
	return i.DeleteItem(itemId, ctx)
}
