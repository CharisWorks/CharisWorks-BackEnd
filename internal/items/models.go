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
	Status Status `json:"status"`
}

type ItemOverview struct {
	Item_id      string                 `json:"item_id"`
	Properties   ItemOverviewProperties `json:"properties"`
	Manufacturer ManufacturerDetails    `json:"manufacturer"`
}

type ManufacturerDetails struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ItemOverviewProperties struct {
	Name    string              `json:"name"`
	Price   int                 `json:"price"`
	Details ItemOverviewDetails `json:"details"`
}

type ItemOverviewDetails struct {
	Status      Status   `json:"status"`
	Stock       int      `json:"stock"`
	Size        int      `json:"size"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
type Status string

const (
	ItemStatusAvailable Status = "Available"
	ItemStatusExpired   Status = "Expired"
	ItemStatusReady     Status = "Ready"
)

type IRequests interface {
	GetOverview(itemId string, itemRepository IRepository) (*ItemOverview, error)
	GetSearchPreviewList(ctx *gin.Context, itemRepository IRepository, itemUtils IUtils) (*[]ItemPreview, int, error)
}

type IRepository interface {
	GetItemOverview(itemId string) (*ItemOverview, error)
	GetPreviewList(pageNum int, pageSize int, conditions map[string]interface{}, tags []string) (*[]ItemPreview, int, error)
}

type IUtils interface {
	InspectSearchConditions(ctx *gin.Context) (pageNum int, pageSize int, conditions map[string]interface{}, tags []string, err error)
}
