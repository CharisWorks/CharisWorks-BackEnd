package items

func ExampleItemPreview() []ItemPreview {
	e := ItemPreview{
		Item_id: "f6d655da-6fff-11ee-b3bc-e86a6465f38b",
		Properties: ItemPreviewProperties{
			Name:  "クラウディ・エンチャント",
			Price: 2480,
			Details: ItemPreviewDetails{
				Status: "Available",
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
				Status:      getStringPointer("Available"),
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

func (i ExampleItemRequests) GetOverview(itemId string) *ItemOverview {
	ItemOverview := ExampleItemOverview(itemId)
	return &ItemOverview
}
func (i ExampleItemRequests) GetPreviewList() *[]ItemPreview {
	ItemPreview := ExampleItemPreview()
	return &ItemPreview
}
func (i ExampleItemRequests) GetSearchPreviewList(tags []string) *[]ItemPreview {
	ItemPreview := ExampleItemPreview()
	return &ItemPreview
}
