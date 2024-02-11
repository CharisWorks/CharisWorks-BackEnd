package items

import (
	"log"
)

func GetOverview(i IItemRequests, itemId string) ItemOverview {
	return *i.GetOverview(itemId)
}

func GetPreviewList(i IItemRequests) []ItemPreview {
	return *i.GetPreviewList()
}

func GetSearchPreviewList(i IItemRequests, keywords []string) []ItemPreview {
	log.Println(keywords)
	return *i.GetPreviewList()
}
