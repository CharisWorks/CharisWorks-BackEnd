package items

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOverview(i IItemRequests, itemId string, ctx *gin.Context) (*ItemOverview, error) {
	return i.GetOverview(itemId, ctx)
}

func GetPreviewList(i IItemRequests, ctx *gin.Context) (*[]ItemPreview, error) {
	item, err := i.GetPreviewList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return nil, err
	}

	return item, nil
}

func GetSearchPreviewList(i IItemRequests, keywords []string, ctx *gin.Context) (*[]ItemPreview, error) {
	log.Println(keywords)
	return i.GetPreviewList(ctx)
}
