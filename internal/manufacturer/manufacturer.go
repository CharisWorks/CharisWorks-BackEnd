package manufacturer

type ManufacturerRequests struct {
}

func (m ManufacturerRequests) RegisterItem(itemRegisterPayload ItemRegisterPayload, userId string, manufacturerDB IRepository, manufacturerUtils IUtils, manufacturerDBHistoy IHistoryRepository) error {
	err := manufacturerUtils.InspectRegisterPayload(itemRegisterPayload)
	if err != nil {
		return err
	}

	err = manufacturerDB.RegisterItem(itemRegisterPayload, userId)
	if err != nil {
		return err
	}
	return nil
}
func (m ManufacturerRequests) UpdateItem(query map[string]interface{}, manufacturerDB IRepository, manufacturerUtils IUtils, manufacturerDBHistoy IHistoryRepository) error {
	updatepayload, err := manufacturerUtils.InspectUpdatePayload(query)
	if err != nil {
		return err
	}

	err = manufacturerDB.UpdateItem(*updatepayload, 1)
	if err != nil {
		return err
	}
	return nil
}
func (m ManufacturerRequests) DeleteItem(itemId string, manufacturerDB IRepository) error {
	err := manufacturerDB.DeleteItem(itemId)
	if err != nil {
		return err
	}
	return nil
}
