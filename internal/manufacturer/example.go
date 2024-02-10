package manufacturer

import "github.com/charisworks/charisworks-backend/internal/items"

type ExampleManufacturerRequests struct {
}

func (m ExampleManufacturerRequests) RegisterItem(i items.ItemOverviewProperties) error {
	return nil
}

func (m ExampleManufacturerRequests) UpdateItem(i items.ItemOverviewProperties) error {
	return nil
}

func (m ExampleManufacturerRequests) DeleteItem(itemId string) error {
	return nil
}
