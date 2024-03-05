package cart

import (
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type Requests struct {
}

func (c Requests) Get(userId string, cartRepository IRepository, cartUtils IUtils) (cart *[]Cart, err error) {
	internalCart, err := cartRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	inspectedCart, err := cartUtils.InspectCart(*internalCart)
	resultCart := cartUtils.ConvertCart(inspectedCart)
	if err != nil {
		return &resultCart, err
	}
	return &resultCart, nil
}

func (c Requests) Register(userId string, cartRequestPayload CartRequestPayload, cartRepository IRepository, cartUtils IUtils) error {
	internalCart, err := cartRepository.Get(userId)
	if err != nil {
		return err
	}

	inspectedCart, _ := cartUtils.InspectCart(*internalCart)
	_, exist := inspectedCart[cartRequestPayload.ItemId]
	itemStatus, err := cartRepository.GetItem(cartRequestPayload.ItemId)
	if err != nil {
		return err
	}
	InspectedCartRequestPayload, err := cartUtils.InspectPayload(cartRequestPayload, *itemStatus)
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

func (c Requests) Delete(userId string, itemId string, cartRepository IRepository, cartUtils IUtils) error {
	internalCart, err := cartRepository.Get(userId)
	if err != nil {
		return err
	}
	inspectedCart, _ := cartUtils.InspectCart(*internalCart)

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
