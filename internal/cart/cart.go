package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartRequests struct {
}

func (c CartRequests) Get(ctx *gin.Context, CartDB ICartDB, CartUtils ICartUtils, userId string) (cart *[]Cart, err error) {
	internalCart, err := CartDB.GetCart(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "cannot get cart"})
		return nil, err
	}
	inspectedCart, err := CartUtils.InspectCart(*internalCart)
	resultCart := CartUtils.ConvertCart(inspectedCart)
	if err != nil {
		ctx.JSON(http.StatusOK, resultCart)
		return nil, err
	}
	return &resultCart, nil
}

func (c CartRequests) Register(CartRequestPayload CartRequestPayload, CartDB ICartDB, CartUtils ICartUtils, ctx *gin.Context, userId string) error {

	internalCart, err := CartDB.GetCart(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "cannot get cart"})
		return err
	}

	inspectedCart, _ := CartUtils.InspectCart(*internalCart)
	_, exist := inspectedCart[CartRequestPayload.ItemId]
	itemStatus, err := CartDB.GetItem(CartRequestPayload.ItemId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
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

func (c CartRequests) Delete(itemId string, CartDB ICartDB, CartUtils ICartUtils, ctx *gin.Context, userId string) error {
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
