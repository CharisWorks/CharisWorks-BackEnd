package items

import (
	"github.com/gin-gonic/gin"
)

func GetOverview(itemId string, i IItemRequests, ctx *gin.Context) (*ItemOverview, error) {
	return i.GetOverview(itemId, ctx)
}

func GetPreviewList(i IItemRequests, ctx *gin.Context) (*[]ItemPreview, error) {

	return i.GetPreviewList(ctx)
}

func GetSearchPreviewList(keywords []string, i IItemRequests, ctx *gin.Context) (*[]ItemPreview, error) {
	return i.GetSearchPreviewList(keywords, ctx)
}
