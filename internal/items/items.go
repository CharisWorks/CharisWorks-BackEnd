package items

import "github.com/gin-gonic/gin"

type Requests struct {
	ItemRepository IRepository
	ItemUtils      IUtils
}

func (r Requests) GetOverview(itemId string) (*Overview, error) {
	return r.ItemRepository.GetItemOverview(itemId)
}
func (r Requests) GetSearchPreviewList(ctx *gin.Context) (*[]Preview, int, error) {
	pageNum, pageSize, inspectedConditions, tags, err := r.ItemUtils.InspectSearchConditions(ctx)
	if err != nil {
		return nil, 0, err
	}
	return r.ItemRepository.GetPreviewList(pageNum, pageSize, inspectedConditions, tags)
}
