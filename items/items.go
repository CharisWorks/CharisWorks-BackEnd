package items

import (
	"log"
)

type ItemRequests struct {
}

func (i ItemRequests) GetOverview(itemId string) *ItemOverview {
	ItemOverview := ExampleItemOverview(itemId)
	return &ItemOverview
}
func (i ItemRequests) GetPreviewList() *[]ItemPreview {
	ItemPreview := ExampleItemPreview()
	return &ItemPreview
}
func (i ItemRequests) GetSearchPreviewList(tags []string) *[]ItemPreview {
	ItemPreview := ExampleItemPreview()
	return &ItemPreview
}

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
