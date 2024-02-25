package items

import "github.com/gin-gonic/gin"

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
	Status ItemStatus `json:"status"`
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
	Status      ItemStatus `json:"status"`
	Stock       *int       `json:"stock"`
	Size        *int       `json:"size"`
	Description *string    `json:"description"`
	Tags        *[]string  `json:"tags"`
}
type ItemStatus string

const (
	ItemStatusAvailable ItemStatus = "Available"
	ItemStatusExpired   ItemStatus = "Expired"
	ItemStatusReady     ItemStatus = "Ready"
)

type IItemRequests interface {
	GetOverview(IItemDB, *gin.Context) (*ItemOverview, error)
	GetSearchPreviewList(ItemDB IItemDB, ItemUtils IItemUtils, ctx *gin.Context) (*[]ItemPreview, error)
}

type IItemDB interface {
	GetItemOverview(itemId string) (*ItemOverview, error)
	GetPreviewList(keywords *[]string, page *string, manufacturer *string) (*[]ItemPreview, error)
}

type IItemUtils interface {
	SortItemsByHighPrice(items *[]ItemPreview) *[]ItemPreview
	SortItemsByLowPrice(items *[]ItemPreview) *[]ItemPreview
	SortItemsByRecommendation(items *[]ItemPreview) *[]ItemPreview
	SortItemsBySize(items *[]ItemPreview) *[]ItemPreview
}
