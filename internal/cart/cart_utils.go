package cart

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type CartUtils struct {
}

func (CartUtils CartUtils) InspectCart(internalCarts []internalCart) (result map[string]internalCart, err error) {
	cartMap := map[string]internalCart{}
	for _, internalCart := range internalCarts {
		if internalCart.itemStock < internalCart.Cart.Quantity {
			internalCart.Cart.ItemProperties.Details.Status = CartItemStatusStockOver
			err = &utils.InternalError{Message: utils.InternalErrorStockOver}
		}
		if internalCart.itemStock == 0 {
			internalCart.Cart.ItemProperties.Details.Status = CartItemStatusNoStock
			err = &utils.InternalError{Message: utils.InternalErrorNoStock}
		}
		if internalCart.status != items.ItemStatusAvailable {
			internalCart.Cart.ItemProperties.Details.Status = CartItemStatusInvalidItem
			err = &utils.InternalError{Message: utils.InternalErrorInvalidItem}
		}
		cartMap[internalCart.Cart.ItemId] = internalCart
	}
	if err != nil {
		return cartMap, &utils.InternalError{Message: utils.InternalErrorInvalidCart}
	}
	return cartMap, nil
}

func (CartUtils CartUtils) ConvertCart(internalCarts map[string]internalCart) (result []Cart) {
	for _, inteinternalCart := range internalCarts {
		Cart := new(Cart)
		Cart = &inteinternalCart.Cart
		result = append(result, *Cart)
	}
	return result
}
func (CartUtils CartUtils) GetTotalAmount(internalCarts map[string]internalCart) int {
	totalAmount := 0
	for _, internalCart := range internalCarts {
		totalAmount += internalCart.Cart.ItemProperties.Price * internalCart.Cart.Quantity
	}
	return totalAmount
}
func (CartUtils CartUtils) InspectPayload(CartRequestPayload CartRequestPayload, itemStatus itemStatus) (result *CartRequestPayload, err error) {
	if CartRequestPayload.Quantity <= 0 {
		return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if itemStatus.status != items.ItemStatusAvailable {
		return nil, &utils.InternalError{Message: utils.InternalErrorInvalidItem}
	}
	if itemStatus.itemStock == 0 {
		return nil, &utils.InternalError{Message: utils.InternalErrorNoStock}
	}
	if CartRequestPayload.Quantity > itemStatus.itemStock {
		return nil, &utils.InternalError{Message: utils.InternalErrorStockOver}
	}

	return &CartRequestPayload, nil
}
