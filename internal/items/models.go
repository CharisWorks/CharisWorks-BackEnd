package items

import "github.com/gin-gonic/gin"

type Preview struct {
	Item_id    string            `json:"item_id"`
	Properties PreviewProperties `json:"properties"`
}

type PreviewProperties struct {
	Name    string         `json:"name"`
	Price   int            `json:"price"`
	Details PreviewDetails `json:"details"`
}

type PreviewDetails struct {
	Status Status `json:"status"`
}

type Overview struct {
	Item_id      string              `json:"item_id"`
	Properties   OverviewProperties  `json:"properties"`
	Manufacturer ManufacturerDetails `json:"manufacturer"`
}

type ManufacturerDetails struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	StripeAccountId string `json:"stripe_account_id"`
	UserId          string
}

type OverviewProperties struct {
	Name    string          `json:"name"`
	Price   int             `json:"price"`
	Details OverviewDetails `json:"details"`
}

type OverviewDetails struct {
	Status      Status   `json:"status"`
	Stock       int      `json:"stock"`
	Size        int      `json:"size"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
type ItemStatus struct {
	Stock  int
	Status Status
}
type Status string

const (
	Available Status = "Available"
	Expired   Status = "Expired"
	Ready     Status = "Ready"
)

type IRequests interface {
	GetOverview(itemId string) (Overview, error)
	GetSearchPreviewList(ctx *gin.Context) ([]Preview, int, error)
}

type IRepository interface {
	GetItemOverview(itemId string) (Overview, error)
	GetPreviewList(pageNum int, pageSize int, conditions map[string]interface{}, tags []string) ([]Preview, int, error)
}

type IUpdater interface {
	ReduceStock(itemId string, Quantity int) error
	StatusUpdate(itemId string, State Status)
}
type IGetStatus interface {
	GetItem(itemId string) (ItemStatus, error)
}
type IUtils interface {
	InspectSearchConditions(ctx *gin.Context) (pageNum int, pageSize int, conditions map[string]interface{}, tags []string, err error)
}
