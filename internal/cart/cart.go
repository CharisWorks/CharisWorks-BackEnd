package cart

import (
	"github.com/gin-gonic/gin"
)

func GetCart(i ICartRequest, ctx *gin.Context) (*[]Cart, error) {
	Cart, err := i.Get(ctx)
	return Cart, err
}
func PostCart(p CartRequestPayload, i ICartRequest, ctx *gin.Context) error {
	return i.Register(p, ctx)
}
func UpdateCart(p CartRequestPayload, i ICartRequest, ctx *gin.Context) error {
	return i.Update(p, ctx)
}
func DeleteCart(itemId string, i ICartRequest, ctx *gin.Context) error {
	return i.Delete(itemId, ctx)
}
