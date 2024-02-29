package items

import (
	"log"
	"strconv"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func ExampleItemPreview() []ItemPreview {
	e := ItemPreview{
		Item_id: 1,
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
func ExampleItemOverview(itemId int) ItemOverview {
	e := ItemOverview{
		Item_id: itemId,
		Properties: &ItemOverviewProperties{
			Name:  "クラウディ・エンチャント",
			Price: 2480,
			Details: ItemOverviewDetails{
				Status:      ItemStatusAvailable,
				Stock:       1,
				Size:        10,
				Description: "foo",
				Tags:        []string{"Golang", "Java"},
			},
		},
	}

	return e
}

type ExampleItemRequests struct {
}

func (i ExampleItemRequests) GetOverview(ItemDB IItemDB, ctx *gin.Context) (*ItemOverview, error) {
	itemId, err := utils.GetParams("itemId", true, ctx)
	if err != nil {
		return nil, err
	}
	log.Println("itemId: ", itemId)
	id, _ := strconv.Atoi(*itemId)
	ItemOverview := ExampleItemOverview(id)
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

func (i ExampleItemDB) GetItemOverview(itemId int) (*ItemOverview, error) {
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
