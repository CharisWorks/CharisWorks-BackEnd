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

type CartRequest struct {
}

func (p CartRequest) Get(ctx *gin.Context) (*[]Cart, error) {
	Cart := ExampleCart()
	return &Cart, nil
}
func (c CartRequest) Register(p CartRequestPayload, ctx *gin.Context) error {
	log.Print(p)
	return nil
}
func (c CartRequest) Update(p CartRequestPayload, ctx *gin.Context) error {
	log.Print(p)
	return nil
}
func (c CartRequest) Delete(itemId string, ctx *gin.Context) error {
	log.Print(itemId)
	return nil
}
