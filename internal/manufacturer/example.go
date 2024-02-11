package manufacturer

type ExampleManufacturerRequests struct {
}

func (m ExampleManufacturerRequests) RegisterItem(i ItemRegisterPayload) error {
	return nil
}

func (m ExampleManufacturerRequests) UpdateItem(i ItemUpdatePayload) error {
	return nil
}

func (m ExampleManufacturerRequests) DeleteItem(itemId string) error {
	return nil
}
