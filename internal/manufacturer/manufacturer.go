package manufacturer

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type Requests struct {
	ManufacturerItemRepository      IItemRepository
	ManufacturerInspectPayloadUtils IInspectPayloadUtils
	ItemRepository                  items.IRepository
}

func (r Requests) Register(itemRegisterPayload RegisterPayload, userId string, itemId string) error {
	err := r.ManufacturerInspectPayloadUtils.Register(itemRegisterPayload)
	if err != nil {
		return err
	}

	err = r.ManufacturerItemRepository.Register(itemId, itemRegisterPayload, userId)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) Update(updatePayload UpdatePayload, userId string, itemId string) error {
	item, err := r.ItemRepository.GetItemOverview(itemId)
	if err != nil {
		return err
	}
	if item.Manufacturer.UserId != userId {
		return &utils.InternalError{Message: utils.InternalErrorInvalidUserRequest}
	}

	updatepayload, err := r.ManufacturerInspectPayloadUtils.Update(updatePayload)
	if err != nil {
		return err
	}
	err = r.ManufacturerItemRepository.Update(updatepayload, itemId)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) Delete(itemId string, userId string) error {
	item, err := r.ItemRepository.GetItemOverview(itemId)
	if err != nil {
		return err
	}
	if item.Manufacturer.StripeAccountId != userId {
		return &utils.InternalError{Message: utils.InternalErrorInvalidUserRequest}
	}
	err = r.ManufacturerItemRepository.Delete(itemId)
	if err != nil {
		return err
	}
	return nil
}
