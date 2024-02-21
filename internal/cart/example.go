package cart

import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func ExampleCart() *[]Cart {
	e := Cart{
		ItemId:   "f6d655da-6fff-11ee-b3bc-e86a6465f38b",
		Quantity: 1,
		ItemProperties: items.ItemPreviewProperties{
			Name:  "クラウディ・エンチャント",
			Price: 2480,
			Details: items.ItemPreviewDetails{
				Status: "Available",
			},
		},
	}
	re := new([]Cart)
	cart := append(*re, e)
	return &cart

}

type ExapleCartRequest struct {
}

func (p ExapleCartRequest) Get(ctx *gin.Context, i ICartDB, userId string) (*[]Cart, error) {
	Cart := ExampleCart()
	return Cart, nil
}
func (c ExapleCartRequest) Register(p CartRequestPayload, i ICartDB, ctx *gin.Context) error {
	log.Print("CartRequestPayload: ", p)
	if p.Quantity <= 0 {
		err := new(utils.InternalError)
		err.SetError("Quantity is invalid")
		return err
	}
	return nil
}
func (c ExapleCartRequest) Delete(itemId string, ctx *gin.Context) error {
	log.Print("itemId: ", itemId)
	return nil
}

type ExampleCartDB struct {
}

func (c ExampleCartDB) GetItem(itemId string) (*int, *string, error) {
	return nil, nil, nil
}
func (c ExampleCartDB) GetCart(userId string) (*[]internalCart, error) {
	return nil, nil
}
func (c ExampleCartDB) RegisterCart(userId string, CartRequestPayload CartRequestPayload) error {
	return nil
}
func (c ExampleCartDB) UpdateCart(userId string, CartRequestPayload CartRequestPayload) error {
	return nil
}
func (c ExampleCartDB) DeleteCart(userId string, itemId string) error {
	return nil
}
