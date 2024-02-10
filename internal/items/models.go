package items

type ItemPreview struct {
	Item_id    string                `json:"item_id"`
	Properties ItemPreviewProperties `json:"properties"`
}

type ItemPreviewProperties struct {
	Name    string             `json:"name"`
	Price   int                `json:"price"`
	Details ItemPreviewDetails `json:"details"`
}

type ItemPreviewDetails struct {
	Status string `json:"status"`
}

type ItemOverview struct {
	Item_id      string                  `json:"item_id"`
	Properties   *ItemOverviewProperties `json:"properties"`
	Manufacturer *ManufacturerDetails    `json:"manufacturer"`
}
type ManufacturerDetails struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ItemOverviewProperties struct {
	Name    *string              `json:"name"`
	Price   *int                 `json:"price"`
	Details *ItemOverviewDetails `json:"details"`
}

type ItemOverviewDetails struct {
	Status      *string   `json:"status"`
	Stock       *int      `json:"stock"`
	Size        *int      `json:"size"`
	Description *string   `json:"description"`
	Tags        *[]string `json:"tags"`
}
type IItemRequests interface {
	GetOverview(itemId string) *ItemOverview
	GetPreviewList() *[]ItemPreview
	GetSearchPreviewList([]string) *[]ItemPreview
}
