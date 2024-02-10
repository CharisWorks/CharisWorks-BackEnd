package manufacturer

import "github.com/charisworks/charisworks-backend/internal/items"

func RegisterItem(p items.ItemOverviewProperties, i IManufacturerRequests) error {
	err := i.RegisterItem(p)
	return err
}
func UpdateItem(p items.ItemOverviewProperties, i IManufacturerRequests) error {
	err := i.UpdateItem(p)
	return err
}
func DeleteItem(itemId string, i IManufacturerRequests) error {
	err := i.DeleteItem(itemId)
	return err
}
