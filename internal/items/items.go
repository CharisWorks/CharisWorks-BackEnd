package items

import (
	"github.com/gin-gonic/gin"
)

func GetOverview(i IItemRequests, itemId string, ctx *gin.Context) (*ItemOverview, error) {
	return i.GetOverview(itemId, ctx)
}

func GetPreviewList(i IItemRequests, ctx *gin.Context) (*[]ItemPreview, error) {

	return i.GetPreviewList(ctx)
}

func GetSearchPreviewList(i IItemRequests, keywords []string, ctx *gin.Context) (*[]ItemPreview, error) {
	return i.GetSearchPreviewList(keywords, ctx)
}
