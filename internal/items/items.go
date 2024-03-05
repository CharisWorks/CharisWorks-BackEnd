package items

import "github.com/gin-gonic/gin"

type ItemRequests struct {
}

func (r ItemRequests) GetOverview(itemId string, itemRepository IRepository) (*ItemOverview, error) {
	return itemRepository.GetItemOverview(itemId)
}
func (r ItemRequests) GetSearchPreviewList(ctx *gin.Context, itemRepository IRepository, itemUtils IUtils) (*[]ItemPreview, int, error) {
	pageNum, pageSize, inspectedConditions, tags, err := itemUtils.InspectSearchConditions(ctx)
	if err != nil {
		return nil, 0, err
	}
	return itemRepository.GetPreviewList(pageNum, pageSize, inspectedConditions, tags)
}
