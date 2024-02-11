package items

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOverview(i IItemRequests, itemId string) (ItemOverview, error) {
	return i.GetOverview(itemId)
}

func GetPreviewList(c *gin.Context, i IItemRequests) (*[]ItemPreview, error) {
	item, err := i.GetPreviewList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil, err
	}

	return &item, nil
}

func GetSearchPreviewList(i IItemRequests, keywords []string) ([]ItemPreview, error) {
	log.Println(keywords)
	return i.GetPreviewList()
}
