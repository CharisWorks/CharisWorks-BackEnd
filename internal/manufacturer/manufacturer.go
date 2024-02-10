package manufacturer

import "github.com/charisworks/charisworks-backend/internal/items"

type ManufacturerRequests struct {
}

func (m ManufacturerRequests) RegisterItem(i items.ItemOverviewProperties) string {
	return ""
}

func (m ManufacturerRequests) UpdateItem(i items.ItemOverviewProperties) string {
	return ""
}

func (m ManufacturerRequests) DeleteItem(itemId string) string {
	return ""
}
