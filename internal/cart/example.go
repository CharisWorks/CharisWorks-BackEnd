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
	itemStatus    *itemStatus
	internalCarts *[]internalCart
	err           error
	updateerr     error
	registererror error
	deleteerror   error
}

func (c ExampleCartDB) GetItem(itemId string) (*itemStatus, error) {
	if c.err != nil {
		return nil, &utils.InternalError{Message: utils.InternalErrorNotFound}
	}
	return c.itemStatus, nil
}
func (c ExampleCartDB) GetCart(userId string) (*[]internalCart, error) {
	if c.internalCarts == nil {
		return nil, &utils.InternalError{Message: utils.InternalErrorNotFound}
	}
	return c.internalCarts, nil
}
func (c ExampleCartDB) RegisterCart(userId string, CartRequestPayload CartRequestPayload) error {
	if c.registererror != nil {
		return c.registererror
	}
	return nil
}
func (c ExampleCartDB) UpdateCart(userId string, CartRequestPayload CartRequestPayload) error {
	if c.updateerr != nil {
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (c ExampleCartDB) DeleteCart(userId string, itemId string) error {
	if c.deleteerror != nil {
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
