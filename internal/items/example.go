package items

/*
import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func ExampleItemPreview() []ItemPreview {
	e := ItemPreview{
		Item_id: "aaa",
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
		Properties: ItemOverviewProperties{
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
	itemId, err := utils.GetParams("itemId", ctx)
	if err != nil {
		return nil, err
	}
	log.Println("itemId: ", itemId)
	ItemOverview := ExampleItemOverview(*itemId)
	return &ItemOverview, nil
}
func (i ExampleItemRequests) GetSearchPreviewList(ItemDb IItemDB, ItemUtils IItemUtils, ctx *gin.Context) (*[]ItemPreview, error) {
	//log.Println("tags: ", tags)
		page, _ := utils.GetQuery("page", false, ctx)
	   	sort, _ := utils.GetQuery("sort", false, ctx)
	   	keywords, _ := utils.GetQuery("keyword", false, ctx)
	   	keywordlist := strings.Split(*keywords, "+")
	   	manufacturer, _ := utils.GetQuery("manufacturer", false, ctx)
	ItemPreview := ExampleItemPreview()
	return &ItemPreview, nil
}
*/
