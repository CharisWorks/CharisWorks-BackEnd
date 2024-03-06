package manufacturer

import (
	"github.com/charisworks/charisworks-backend/internal/items"
)

type Requests struct {
	ManufacturerItemRepository      IItemRepository
	ManufacturerInspectPayloadUtils IInspectPayloadUtils
	ItemRepository                  items.IRepository
}

func (r Requests) Register(itemRegisterPayload RegisterPayload, userId string) error {
	err := r.ManufacturerInspectPayloadUtils.Register(itemRegisterPayload)
	if err != nil {
		return err
	}

	err = r.ManufacturerItemRepository.Register("", itemRegisterPayload, userId)
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) Update(updatePayload UpdatePayload, userId string) error {
	updatepayload, err := r.ManufacturerInspectPayloadUtils.Update(updatePayload)
	if err != nil {
		return err
	}

	err = r.ManufacturerItemRepository.Update(updatepayload, "")
	if err != nil {
		return err
	}
	return nil
}
func (r Requests) Delete(itemId string, userId string) error {
	err := r.ManufacturerItemRepository.Delete(itemId)
	if err != nil {
		return err
	}
	return nil
}
