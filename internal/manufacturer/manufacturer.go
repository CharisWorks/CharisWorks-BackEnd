package manufacturer

func RegisterItem(p ItemRegisterPayload, i IManufacturerRequests) error {
	err := i.RegisterItem(p)
	return err
}
func UpdateItem(p ItemUpdatePayload, i IManufacturerRequests) error {
	err := i.UpdateItem(p)
	return err
}
func DeleteItem(itemId string, i IManufacturerRequests) error {
	err := i.DeleteItem(itemId)
	return err
}
