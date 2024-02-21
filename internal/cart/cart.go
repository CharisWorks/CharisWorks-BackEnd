package cart

import (
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type CartRequest struct {
}

func (c CartRequest) Get(ctx *gin.Context, CartDB ICartDB, CartUtils ICartUtils, userId string) (*[]Cart, error) {
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
	errorList := new([]string)
	cartMap := map[string]Cart{}
	for _, internalCart := range internalCarts {
		if internalCart.itemStock < internalCart.Cart.Quantity {
			internalCart.Cart.ItemProperties.Details.Status = "stock over"
			err = &utils.InternalError{Message: "stock over"}
		}
		if internalCart.itemStock == 0 {
			internalCart.Cart.ItemProperties.Details.Status = "no stock"
			err = &utils.InternalError{Message: "no stock"}
		}
		if internalCart.status != "Available" {
			internalCart.Cart.ItemProperties.Details.Status = internalCart.status
			err = &utils.InternalError{Message: "invalid item"}
		}
		if err != nil {
			*errorList = append(*errorList, internalCart.Cart.ItemId+": "+err.Error())
		}
		err = nil
		cartMap[internalCart.Cart.ItemId] = internalCart.Cart
	}
	if len(*errorList) != 0 {
		mes := new(string)
		for _, e := range *errorList {
			*mes += e + "\n"
		}
		return result, &utils.InternalError{
			Message: *mes,
		}
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
	if itemStatus.status != "Available" {
		return nil, &utils.InternalError{Message: "invalid item"}
	}
	if CartRequestPayload.Quantity > itemStatus.itemStock {
		return nil, &utils.InternalError{Message: "stock over"}
	}
	if CartRequestPayload.Quantity == 0 {
		return nil, &utils.InternalError{Message: "invalid quantity"}
	}
	return &CartRequestPayload, nil
}
