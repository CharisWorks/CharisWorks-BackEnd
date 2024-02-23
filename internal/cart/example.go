package cart

import (
	"github.com/charisworks/charisworks-backend/internal/utils"
)

/*
	func ExampleCart() *[]Cart {
		e := Cart{
			ItemId:   "f6d655da-6fff-11ee-b3bc-e86a6465f38b",
			Quantity: 1,
			ItemProperties: CartItemPreviewProperties{
				Name:  "クラウディ・エンチャント",
				Price: 2000,
				Details: CartItemPreviewDetails{
					Status: CartItemStatusAvailable,
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
			return &utils.InternalError{Message: utils.InternalErrorInvalidQuantity}
		}
		return nil
	}

	func (c ExapleCartRequest) Delete(itemId string, ctx *gin.Context) error {
		log.Print("itemId: ", itemId)
		return nil
	}

	func (c ExapleCartRequest) GetItem(itemid string) (*itemStatus, error) {
		return nil, nil
	}
*/
type ExampleCartDB struct {
	ItemStatus    *itemStatus
	InternalCarts *[]InternalCart
	Err           error
	UpdateErr     error
	RegisterError error
	DeleteError   error
}

func (c ExampleCartDB) GetItem(itemId string) (*itemStatus, error) {
	if c.Err != nil {
		return nil, &utils.InternalError{Message: utils.InternalErrorNotFound}
	}
	return c.ItemStatus, nil
}
func (c ExampleCartDB) GetCart(userId string) (*[]InternalCart, error) {
	if c.InternalCarts == nil {
		return nil, &utils.InternalError{Message: utils.InternalErrorNotFound}
	}
	return c.InternalCarts, nil
}
func (c ExampleCartDB) RegisterCart(userId string, CartRequestPayload CartRequestPayload) error {
	if c.RegisterError != nil {
		return c.RegisterError
	}
	return nil
}
func (c ExampleCartDB) UpdateCart(userId string, CartRequestPayload CartRequestPayload) error {
	if c.UpdateErr != nil {
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (c ExampleCartDB) DeleteCart(userId string, itemId string) error {
	if c.DeleteError != nil {
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
