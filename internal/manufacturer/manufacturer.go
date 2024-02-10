package manufacturer

import "github.com/charisworks/charisworks-backend/internal/items"

func RegisterItem(p items.ItemOverviewProperties, i IManufacturerRequests) error {
	err := i.RegisterItem(p)
	return err
}
