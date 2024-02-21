package cart

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type CartRequest struct {
}

func (c CartRequest) Get(ctx *gin.Context, CartDB ICartDB, CartUtils ICartUtils, userId string) (cart *[]Cart, err error) {
	resultCart := new([]Cart)
	internalCart, err := CartDB.GetCart(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "cannot get cart"})
		return nil, err
	}
	inspectedCart, err := CartUtils.InspectCart(*internalCart)
	resultCart = CartUtils.ConvertCart(inspectedCart)
	if err != nil {
		ctx.JSON(http.StatusOK, resultCart)
		return nil, err
	}
	return resultCart, nil
}

func (c CartRequest) Register(CartRequestPayload CartRequestPayload, CartDB ICartDB, CartUtils ICartUtils, ctx *gin.Context, userId string) error {
	internalCart, err := CartDB.GetCart(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "cannot get cart"})
		return err
	}
	inspectedCart, _ := CartUtils.InspectCart(*internalCart)
	_, exist := inspectedCart[CartRequestPayload.ItemId]
	itemStatus, err := CartDB.GetItem(CartRequestPayload.ItemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return err
	}
	InspectedCartRequestPayload, err := CartUtils.InspectPayload(CartRequestPayload, *itemStatus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}
	if exist {
		err = CartDB.UpdateCart(userId, *InspectedCartRequestPayload)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return err
		}
	} else {
		err = CartDB.RegisterCart(userId, *InspectedCartRequestPayload)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return err
		}
	}
	return nil
}

func (c CartRequest) Delete(itemId string, CartDB ICartDB, CartUtils ICartUtils, ctx *gin.Context, userId string) error {
	internalCart, err := CartDB.GetCart(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "cannot get cart"})
		return err
	}
	inspectedCart, _ := CartUtils.InspectCart(*internalCart)
	_, exist := inspectedCart[itemId]
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "this item is not exist in cart"})
		return err
	}
	err = CartDB.DeleteCart(userId, itemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return err
	}
	return nil
}

type CartUtils struct {
}

func (CartUtils CartUtils) InspectCart(internalCarts []internalCart) (result map[string]internalCart, err error) {

	cartMap := map[string]Cart{}
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
		err = nil
		cartMap[internalCart.Cart.ItemId] = internalCart.Cart
	}
	if err != nil {
		return result, &utils.InternalError{Message: utils.InternalErrorInvalidCart}
	}
	return result, nil
}

func (CartUtils CartUtils) ConvertCart(internalCarts map[string]internalCart) (result *[]Cart) {
	for _, inteinternalCart := range internalCarts {
		Cart := new(Cart)
		Cart = &inteinternalCart.Cart
		*result = append(*result, *Cart)
	}
	return result
}
func (CartUtils CartUtils) GetTotalAmount(internalCarts map[string]internalCart) int {
	totalAmount := 0
	for _, internalCart := range internalCarts {
		totalAmount += internalCart.Cart.ItemProperties.Price
	}
	return totalAmount
}
func (CartUtils CartUtils) InspectPayload(CartRequestPayload CartRequestPayload, itemStatus itemStatus) (result *CartRequestPayload, err error) {
	if itemStatus.status != items.ItemStatusAvailable {
		return nil, &utils.InternalError{Message: utils.InternalErrorInvalidItem}
	}
	if CartRequestPayload.Quantity > itemStatus.itemStock {
		return nil, &utils.InternalError{Message: utils.InternalErrorStockOver}
	}
	if CartRequestPayload.Quantity == 0 {
		return nil, &utils.InternalError{Message: utils.InternalErrorNoStock}
	}
	return &CartRequestPayload, nil
}
