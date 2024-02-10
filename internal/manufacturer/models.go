package manufacturer

import "github.com/charisworks/charisworks-backend/internal/items"

type IManufacturerRequests interface {
	RegisterItem(i items.ItemOverviewProperties) error
	UpdateItem(i items.ItemOverviewProperties) error
	DeleteItem(itemId string) error
}
