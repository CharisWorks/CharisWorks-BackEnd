package cart

import (
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type CartRequests struct {
}

func (c CartRequests) Get(userId string, cartRepository ICartRepository, CartUtils ICartUtils) (cart *[]Cart, err error) {
	internalCart, err := cartRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	inspectedCart, err := CartUtils.Inspect(*internalCart)
	resultCart := CartUtils.Convert(inspectedCart)
	if err != nil {
		return &resultCart, err
	}
	return &resultCart, nil
}

func (c CartRequests) Register(userId string, cartRequestPayload CartRequestPayload, cartRepository ICartRepository, CartUtils ICartUtils) error {
	internalCart, err := cartRepository.Get(userId)
	if err != nil {
		return err
	}

	inspectedCart, _ := CartUtils.Inspect(*internalCart)
	_, exist := inspectedCart[cartRequestPayload.ItemId]
	itemStatus, err := cartRepository.GetItem(cartRequestPayload.ItemId)
	if err != nil {
		return err
	}
	InspectedCartRequestPayload, err := CartUtils.InspectPayload(cartRequestPayload, *itemStatus)
	if err != nil {
		return err
	}
	if exist {
		err = cartRepository.Update(userId, *InspectedCartRequestPayload)
		if err != nil {
			return err
		}
	} else {
		err = cartRepository.Register(userId, *InspectedCartRequestPayload)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c CartRequests) Delete(userId string, itemId string, cartRepository ICartRepository, CartUtils ICartUtils) error {
	internalCart, err := cartRepository.Get(userId)
	if err != nil {
		return err
	}
	inspectedCart, _ := CartUtils.Inspect(*internalCart)

	_, exist := inspectedCart[itemId]
	if !exist {
		return &utils.InternalError{Message: utils.InternalErrorInvalidUserRequest}
	}
	err = cartRepository.Delete(userId, itemId)
	if err != nil {
		return err
	}
	return nil
}
