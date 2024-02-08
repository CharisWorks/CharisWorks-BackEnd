package items

func ExampleItemPreview() []ItemPreview {
	e := new(ItemPreview)
	e.Item_id = "f6d655da-6fff-11ee-b3bc-e86a6465f38b"
	e.Properties.Name = "クラウディ・エンチャント"
	e.Properties.Price = 2480
	e.Properties.Details.Status = "Available"
	re := new([]ItemPreview)
	return append(*re, *e)

}
func ExampleItemOverview(itemId string) ItemOverview {
	e := new(ItemOverview)
	e.Item_id = itemId
	e.Properties.Name = "クラウディ・エンチャント"
	e.Properties.Price = 2480
	e.Properties.Details.Status = "Available"
	e.Properties.Details.Stock = 1
	e.Properties.Details.Size = 10
	e.Properties.Details.Description = "foo"
	e.Properties.Details.Tags = []string{"Golang", "Java"}

	return *e

}
