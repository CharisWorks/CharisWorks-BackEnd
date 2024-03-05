package items

import "github.com/gin-gonic/gin"

type ItemRequests struct {
}

func (r ItemRequests) GetOverview(itemId string, ItemRepository IItemRepository) (*Overview, error) {
	return ItemRepository.GetItemOverview(itemId)
}
func (r ItemRequests) GetSearchPreviewList(ctx *gin.Context, ItemRepository IItemRepository, ItemUtils IItemUtils) (*[]Preview, int, error) {
	pageNum, pageSize, inspectedConditions, tags, err := ItemUtils.InspectSearchConditions(ctx)
	if err != nil {
		return nil, 0, err
	}
	return ItemRepository.GetPreviewList(pageNum, pageSize, inspectedConditions, tags)
}
