package items

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ItemRequests struct {
}

func (i ItemRequests) GetOverview(itemId string) ItemOverview {
	return ExampleItemOverview(itemId)
}
func (i ItemRequests) GetPreviewList() []ItemPreview {
	return ExampleItemPreview()
}
func (i ItemRequests) GetSearchPreviewList(tags []string) []ItemPreview {
	return ExampleItemPreview()
}

func GetOverview(c *gin.Context, i IItemRequests, itemId string) ItemOverview {
	return i.GetOverview(itemId)
}

func GetPreviewList(c *gin.Context, i IItemRequests) []ItemPreview {
	return i.GetPreviewList()
}

func GetSearchPreviewList(c *gin.Context, i IItemRequests, keywords []string) []ItemPreview {
	log.Println(keywords)
	return i.GetPreviewList()
}
