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
		Properties: ItemOverviewProperties{
			Name:  "クラウディ・エンチャント",
			Price: 2480,
			Details: ItemOverviewDetails{
				Status:      "Available",
				Stock:       1,
				Size:        10,
				Description: "foo",
				Tags:        []string{"Golang", "Java"},
			},
		},
	}

	return e
}
