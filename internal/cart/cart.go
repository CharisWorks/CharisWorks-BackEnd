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
	internalCart, err = CartUtils.InspectCart(*internalCart)
	resultCart = CartUtils.ConvertCart(*internalCart)
	if err != nil {
		ctx.JSON(http.StatusOK, resultCart)
		return nil, err
	}
	return resultCart, nil
}

func (c CartRequest) Register(CartRequestPayload CartRequestPayload, CartDB ICartDB, ctx *gin.Context, userId string) error {
	internalCart, err := CartDB.GetCart(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "cannot get cart"})
		return err
	}
	for _, cart := range *internalCart {
		if CartRequestPayload.ItemId == cart.Cart.ItemId {
			if cart.itemStock == 0 {
				err = CartDB.DeleteCart(userId, CartRequestPayload.ItemId)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, err)
					return err
				}
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "out of stock"})
				err = &utils.InternalError{Message: "out of stock"}
			}
			if cart.itemStock < CartRequestPayload.Quantity {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "stock over"})
				err = &utils.InternalError{Message: "stock over"}

			}
			if cart.itemStock < cart.Cart.Quantity {
				CartRequestPayload.Quantity = cart.itemStock
				err = CartDB.UpdateCart(userId, CartRequestPayload)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, err)
					return err
				}
				err = &utils.InternalError{Message: "quantity of item from cart is out of stock"}
			}
			err = CartDB.UpdateCart(userId, CartRequestPayload)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return err
			}
		}
	}
	return nil
}

type CartUtils struct {
}

func (CartUtils CartUtils) InspectCart(internalCarts []internalCart) (result *[]internalCart, err error) {
	errorList := new([]string)
	for _, internalCart := range internalCarts {
		if internalCart.itemStock < internalCart.Cart.Quantity {
			internalCart.Cart.ItemProperties.Details.Status = "stock over"
			err = &utils.InternalError{Message: "stock over"}
		}
		if internalCart.itemStock == 0 {
			internalCart.Cart.ItemProperties.Details.Status = "no stock"
			err = &utils.InternalError{Message: "no stock"}
		}
		if internalCart.status == "Available" {
			internalCart.Cart.ItemProperties.Details.Status = internalCart.status
			err = &utils.InternalError{Message: "invalid item"}
		}
		if err != nil {
			*errorList = append(*errorList, err.Error()+": "+internalCart.Cart.ItemId)
		}
		err = nil
		*result = append(*result, internalCart)
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

func (CartUtils CartUtils) ConvertCart(internalCarts []internalCart) (result *[]Cart) {
	for _, inteinternalCart := range internalCarts {
		Cart := new(Cart)
		Cart = &inteinternalCart.Cart
		*result = append(*result, *Cart)
	}
	return result
}
