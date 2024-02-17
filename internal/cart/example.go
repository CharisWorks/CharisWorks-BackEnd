package cart

import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

func ExampleCart() []Cart {
	e := Cart{
		ItemId:   "f6d655da-6fff-11ee-b3bc-e86a6465f38b",
		Quantity: 1,
		ItemPreviewProperties: items.ItemPreviewProperties{
			Name:  "クラウディ・エンチャント",
			Price: 2480,
			Details: items.ItemPreviewDetails{
				Status: "Available",
			},
		},
	}
	re := new([]Cart)
	return append(*re, e)

}

type ExapleCartRequest struct {
}

func (p ExapleCartRequest) Get(ctx *gin.Context) (*[]Cart, error) {
	Cart := ExampleCart()
	return &Cart, nil
}
func (c ExapleCartRequest) Register(p CartRequestPayload, ctx *gin.Context) error {
	log.Print("CartRequestPayload: ", p)
	return nil
}
func (c ExapleCartRequest) Update(p CartRequestPayload, ctx *gin.Context) error {
	log.Print("CartRequestPayload: ", p)
	return nil
}
func (c ExapleCartRequest) Delete(itemId string, ctx *gin.Context) error {
	log.Print("itemId: ", itemId)
	return nil
}
