package manufacturer

type Requests struct {
	ManufacturerItemRepository      IItemRepository
	ManufacturerInspectPayloadUtils IInspectPayloadUtils
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
func (r Requests) Update(updatePayload UpdatePayload) error {
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
func (r Requests) Delete(itemId string) error {
	err := r.ManufacturerItemRepository.Delete(itemId)
	if err != nil {
		return err
	}
	return nil
}
