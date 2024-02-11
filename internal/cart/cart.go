package cart

import (
	"github.com/gin-gonic/gin"
)

func GetCart(i ICartRequest, ctx *gin.Context) (*[]Cart, error) {
	Cart, err := i.Get(ctx)
	return Cart, err
}
func PostCart(p CartRequestPayload, i ICartRequest, ctx *gin.Context) error {
	err := i.Register(p, ctx)
	if err != nil {
		return err
	}
	return err
}
func UpdateCart(p CartRequestPayload, i ICartRequest, ctx *gin.Context) error {
	err := i.Update(p, ctx)
	return err
}
func DeleteCart(itemId string, i ICartRequest, ctx *gin.Context) error {
	err := i.Delete(itemId, ctx)
	return err
}
