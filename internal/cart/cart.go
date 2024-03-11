package cart

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type Requests struct {
	CartRepository IRepository
	CartUtils      IUtils
	ItemGetStatus  items.IGetStatus
}

func (r Requests) Get(userId string) (cart *[]Cart, err error) {
	internalCart, err := r.CartRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	inspectedCart, err := r.CartUtils.Inspect(*internalCart)
	resultCart := r.CartUtils.Convert(inspectedCart)
	if err != nil {
		return &resultCart, err
	}
	return &resultCart, nil
}

func (r Requests) Register(userId string, cartRequestPayload CartRequestPayload) error {
	internalCart, err := r.CartRepository.Get(userId)
	if err != nil {
		return err
	}

	inspectedCart, _ := r.CartUtils.Inspect(*internalCart)
	_, exist := inspectedCart[cartRequestPayload.ItemId]
	itemStatus, err := r.ItemGetStatus.GetItem(cartRequestPayload.ItemId)
	if err != nil {
		return err
	}
	InspectedCartRequestPayload, err := r.CartUtils.InspectPayload(cartRequestPayload, itemStatus)
	if err != nil {
		return err
	}
	if exist {
		err = r.CartRepository.Update(userId, *InspectedCartRequestPayload)
		if err != nil {
			return err
		}
	} else {
		err = r.CartRepository.Register(userId, *InspectedCartRequestPayload)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r Requests) Delete(userId string, itemId string) error {
	internalCart, err := r.CartRepository.Get(userId)
	if err != nil {
		return err
	}
	inspectedCart, _ := r.CartUtils.Inspect(*internalCart)

	_, exist := inspectedCart[itemId]
	if !exist {
		return &utils.InternalError{Message: utils.InternalErrorInvalidUserRequest}
	}
	err = r.CartRepository.Delete(userId, itemId)
	if err != nil {
		return err
	}
	return nil
}
