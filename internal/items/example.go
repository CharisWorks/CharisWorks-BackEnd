package items

import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func ExampleItemPreview() []ItemPreview {
	e := ItemPreview{
		Item_id: "f6d655da-6fff-11ee-b3bc-e86a6465f38b",
		Properties: ItemPreviewProperties{
			Name:  "クラウディ・エンチャント",
			Price: 2480,
			Details: ItemPreviewDetails{
				Status: ItemStatusAvailable,
			},
		},
	}

	re := new([]ItemPreview)
	return append(*re, e)

}
func ExampleItemOverview(itemId string) ItemOverview {
	e := ItemOverview{
		Item_id: itemId,
		Properties: &ItemOverviewProperties{
			Name:  getStringPointer("クラウディ・エンチャント"),
			Price: getIntPointer(2480),
			Details: &ItemOverviewDetails{
				Status:      ItemStatusAvailable,
				Stock:       getIntPointer(1),
				Size:        getIntPointer(10),
				Description: getStringPointer("foo"),
				Tags:        &[]string{"Golang", "Java"},
			},
		},
	}

	return e
}
func getStringPointer(s string) *string {
	return &s
}

func getIntPointer(i int) *int {
	return &i
}

type ExampleItemRequests struct {
}

func (i ExampleItemRequests) GetOverview(ItemDB IItemDB, ctx *gin.Context) (*ItemOverview, error) {
	itemId, err := utils.GetParams("itemId", true, ctx)
	if err != nil {
		return nil, err
	}
	log.Println("itemId: ", itemId)
	ItemOverview := ExampleItemOverview(*itemId)
	return &ItemOverview, nil
}
func (i ExampleItemRequests) GetSearchPreviewList(ItemDb IItemDB, ItemUtils IItemUtils, ctx *gin.Context) (*[]ItemPreview, error) {
	//log.Println("tags: ", tags)
	/* 	page, _ := utils.GetQuery("page", false, ctx)
	   	sort, _ := utils.GetQuery("sort", false, ctx)
	   	keywords, _ := utils.GetQuery("keyword", false, ctx)
	   	keywordlist := strings.Split(*keywords, "+")
	   	manufacturer, _ := utils.GetQuery("manufacturer", false, ctx) */
	ItemPreview := ExampleItemPreview()
	return &ItemPreview, nil
}

type ExampleItemDB struct {
}

func (i ExampleItemDB) GetItemOverview(itemId string) (*ItemOverview, error) {
	ItemOverview := ExampleItemOverview(itemId)
	return &ItemOverview, nil
}
func (i ExampleItemDB) GetPreviewList(keywords *[]string, page *string, manufacturer *string) (*[]ItemPreview, error) {
	ItemPreview := ExampleItemPreview()
	return &ItemPreview, nil
}

type ExampleItemUtils struct {
}

func (i ExampleItemUtils) SortItemsByHighPrice(items *[]ItemPreview) *[]ItemPreview {
	return items
}
func (i ExampleItemUtils) SortItemsByLowPrice(items *[]ItemPreview) *[]ItemPreview {
	return items
}
func (i ExampleItemUtils) SortItemsByRecommendation(items *[]ItemPreview) *[]ItemPreview {
	return items
}
func (i ExampleItemUtils) SortItemsBySize(items *[]ItemPreview) *[]ItemPreview {
	return items
}
