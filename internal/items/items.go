package items

import "github.com/gin-gonic/gin"

type ItemRequests struct {
}

func (r ItemRequests) GetOverview(itemId string, ItemDB IItemDB) (*ItemOverview, error) {
	return ItemDB.GetItemOverview(itemId)
}
func (r ItemRequests) GetSearchPreviewList(ctx *gin.Context, ItemDB IItemDB, ItemUtils IItemUtils) (*[]ItemPreview, error) {
	pageNum, pageSize, inspectedConditions, tags, err := ItemUtils.InspectSearchConditions(ctx)
	if err != nil {
		return nil, err
	}
	return ItemDB.GetPreviewList(pageNum, pageSize, inspectedConditions, tags)
}
