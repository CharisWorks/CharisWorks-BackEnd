package cart

import (
	"github.com/charisworks/charisworks-backend/items"
)

func ExampleCart() []Cart {
	e := Cart{
		ItemId: "f6d655da-6fff-11ee-b3bc-e86a6465f38b",
		ItemPreviewProperties: items.ItemPreviewProperties{
			Name:  "クラウディ・エンチャント",
			Price: 2480,
			Details: items.ItemPreviewDetails{
				Status: "Available",
			},
		},
	}
	re := new([]Cart)
	return append(*re, e)

}