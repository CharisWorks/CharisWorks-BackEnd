package items

import (
	"log"
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

func GetOverview(i IItemRequests, itemId string) ItemOverview {
	return i.GetOverview(itemId)
}

func GetPreviewList(i IItemRequests) []ItemPreview {
	return i.GetPreviewList()
}

func GetSearchPreviewList(i IItemRequests, keywords []string) []ItemPreview {
	log.Println(keywords)
	return i.GetPreviewList()
}
