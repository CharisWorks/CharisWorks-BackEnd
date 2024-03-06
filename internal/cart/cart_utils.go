package cart

import (
	"log"
	"sort"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type Utils struct {
}

func (r Utils) Inspect(internalCarts []InternalCart) (result map[string]InternalCart, err error) {
	cartMap := map[string]InternalCart{}
	for _, internalCart := range internalCarts {
		if internalCart.ItemStock < internalCart.Cart.Quantity {
			internalCart.Cart.ItemProperties.Details.Status = StockOver
			err = &utils.InternalError{Message: utils.InternalErrorStockOver}
		}
		if internalCart.ItemStock == 0 {
			internalCart.Cart.ItemProperties.Details.Status = NoStock
			err = &utils.InternalError{Message: utils.InternalErrorNoStock}
		}
		if internalCart.Status != items.Available {
			internalCart.Cart.ItemProperties.Details.Status = InvalidItem
			err = &utils.InternalError{Message: utils.InternalErrorInvalidItem}
		}
		cartMap[internalCart.Cart.ItemId] = internalCart
	}
	if err != nil {
		return cartMap, &utils.InternalError{Message: utils.InternalErrorInvalidCart}
	}
	return cartMap, nil
}

func (r Utils) Convert(internalCarts map[string]InternalCart) (result []Cart) {
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
func (r Utils) GetTotalAmount(internalCarts map[string]InternalCart) int {
	totalAmount := 0
	for _, internalCart := range internalCarts {
		totalAmount += internalCart.Cart.ItemProperties.Price * internalCart.Cart.Quantity
	}
	return totalAmount
}
func (r Utils) InspectPayload(CartRequestPayload CartRequestPayload, itemStatus itemStatus) (result *CartRequestPayload, err error) {
	if CartRequestPayload.Quantity <= 0 {
		return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if itemStatus.status != items.Available {
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
