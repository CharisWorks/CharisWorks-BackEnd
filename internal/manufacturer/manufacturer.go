package manufacturer

type Requests struct {
}

func (m Requests) Register(itemRegisterPayload ItemRegisterPayload, userId string, manufacturerDB IItemRepository, manufacturerUtils IInspectPayloadUtils, manufacturerDBHistoy IHistoryRepository) error {
	err := manufacturerUtils.Register(itemRegisterPayload)
	if err != nil {
		return err
	}

	err = manufacturerDB.Register("", itemRegisterPayload, userId)
	if err != nil {
		return err
	}
	return nil
}
func (m Requests) Update(query map[string]interface{}, manufacturerDB IItemRepository, manufacturerUtils IInspectPayloadUtils, manufacturerDBHistoy IHistoryRepository) error {
	updatepayload, err := manufacturerUtils.Update(query)
	if err != nil {
		return err
	}

	err = manufacturerDB.Update(*updatepayload, "")
	if err != nil {
		return err
	}
	return nil
}
func (m Requests) Delete(itemId string, manufacturerDB IItemRepository) error {
	err := manufacturerDB.Delete(itemId)
	if err != nil {
		return err
	}
	return nil
}
