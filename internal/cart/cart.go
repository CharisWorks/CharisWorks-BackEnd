package cart

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type CartRequests struct {
}

func (c CartRequests) Get(ctx *gin.Context, CartDB ICartDB, CartUtils ICartUtils) (cart *[]Cart, err error) {

	internalCart, err := CartDB.GetCart(ctx.MustGet("UserId").(string))
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

func (c CartRequests) Register(CartDB ICartDB, CartUtils ICartUtils, ctx *gin.Context) error {
	UserId := ctx.MustGet("UserId").(string)
	payload, err := utils.GetPayloadFromBody(ctx, &CartRequestPayload{})
	if err != nil {
		return err
	}
	CartRequestPayload := new(CartRequestPayload)
	*CartRequestPayload = *payload
	internalCart, err := CartDB.GetCart(UserId)
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
	InspectedCartRequestPayload, err := CartUtils.InspectPayload(*CartRequestPayload, *itemStatus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}
	if exist {
		err = CartDB.UpdateCart(UserId, *InspectedCartRequestPayload)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return err
		}
	} else {
		err = CartDB.RegisterCart(UserId, *InspectedCartRequestPayload)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return err
		}
	}
	return nil
}

func (c CartRequests) Delete(CartDB ICartDB, CartUtils ICartUtils, ctx *gin.Context) error {
	log.Print(ctx.Request.URL.Query())
	itemId, err := utils.GetQuery("item_id", true, ctx)
	if err != nil {
		return err
	}
	UserId := ctx.MustGet("UserId").(string)
	internalCart, err := CartDB.GetCart(UserId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "cannot get cart"})
		return err
	}
	inspectedCart, _ := CartUtils.InspectCart(*internalCart)

	_, exist := inspectedCart[*itemId]
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "this item is not exist in cart"})
		return err
	}
	err = CartDB.DeleteCart(UserId, *itemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return err
	}
	return nil
}
