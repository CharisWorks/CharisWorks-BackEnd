package cart

import (
	"log"
	"sort"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type Utils struct {
}

func (u Utils) InspectCart(internalCarts []InternalCart) (result map[string]InternalCart, err error) {
	cartMap := map[string]InternalCart{}
	for _, internalCart := range internalCarts {
		if internalCart.ItemStock < internalCart.Cart.Quantity {
			internalCart.Cart.ItemProperties.Details.Status = CartItemStatusStockOver
			err = &utils.InternalError{Message: utils.InternalErrorStockOver}
		}
		if internalCart.ItemStock == 0 {
			internalCart.Cart.ItemProperties.Details.Status = CartItemStatusNoStock
			err = &utils.InternalError{Message: utils.InternalErrorNoStock}
		}
		if internalCart.Status != items.ItemStatusAvailable {
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

func (u Utils) ConvertCart(internalCarts map[string]InternalCart) (result []Cart) {
	sortedCart := []InternalCart{}
	for _, inteinternalCart := range internalCarts {
		sortedCart = append(sortedCart, inteinternalCart)
	}
	sort.Slice(sortedCart, func(i, j int) bool { return sortedCart[i].Index < sortedCart[j].Index })
	log.Print("sortedCart: ", sortedCart)
	for _, inteinternalCart := range sortedCart {
		Cart := new(Cart)
		Cart = &inteinternalCart.Cart
		result = append(result, *Cart)
	}

	return result
}
func (u Utils) GetTotalAmount(internalCarts map[string]InternalCart) int {
	totalAmount := 0
	for _, internalCart := range internalCarts {
		totalAmount += internalCart.Cart.ItemProperties.Price * internalCart.Cart.Quantity
	}
	return totalAmount
}
func (u Utils) InspectPayload(cartRequestPayload CartRequestPayload, itemStatus itemStatus) (result *CartRequestPayload, err error) {
	if cartRequestPayload.Quantity <= 0 {
		return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if itemStatus.status != items.ItemStatusAvailable {
		return nil, &utils.InternalError{Message: utils.InternalErrorInvalidItem}
	}
	if itemStatus.itemStock == 0 {
		return nil, &utils.InternalError{Message: utils.InternalErrorNoStock}
	}
	if cartRequestPayload.Quantity > itemStatus.itemStock {
		return nil, &utils.InternalError{Message: utils.InternalErrorStockOver}
	}

	return &cartRequestPayload, nil
}
