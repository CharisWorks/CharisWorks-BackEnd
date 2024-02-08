package items

import "google.golang.org/genproto/googleapis/type/decimal"

type ItemPreview struct {
	Item_id        string                `json:"item_id"`
	ItemProperties ItemPreviewProperties `json:"properties"`
}

type ItemPreviewProperties struct {
	Name    string             `json:"name"`
	Price   decimal.Decimal    `json:"price"`
	Details ItemPreviewDetails `json:"details"`
}

type ItemPreviewDetails struct {
	Status string `json:"status"`
}

type ItemOveview struct {
	Item_id      string                 `json:"item_id"`
	Properties   ItemOverviewProperties `json:"properties"`
	Manufacturer Manufacturer           `json:"manufacturer"`
}

type ItemOverviewProperties struct {
	Name    string              `json:"name"`
	Price   decimal.Decimal     `json:"price"`
	Details ItemOverviewDetails `json:"details"`
}

type ItemOverviewDetails struct {
	Status      string   `json:"status"`
	Stock       int      `json:"stock"`
	Size        int      `json:"size"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
