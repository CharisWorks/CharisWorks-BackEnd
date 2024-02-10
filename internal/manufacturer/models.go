package manufacturer

import "github.com/charisworks/charisworks-backend/internal/items"

type IManufacturerRequests interface {
	RegisterItem(i items.ItemOverviewProperties) (message string)
	UpdateItem(i items.ItemOverviewProperties) (message string)
	DeleteItem(itemId string) (message string)
}
