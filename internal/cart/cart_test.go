package cart

import (
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

func TestGET(t *testing.T) {
	e := new(ExapleCartRequest)
	ctx := new(gin.Context)
	_, err := e.Get(ctx, ExampleCartDB{}, "")
	if err != nil {
		t.Errorf("error")
	}
}
func TestCartRegister(t *testing.T) {
	e := new(ExapleCartRequest)
	ctx := new(gin.Context)
	goodCases := []struct {
		name     string
		ItemId   string
		Quantity int
	}{
		{
			"正常なパターン", "a", 1,
		},
	}
	for _, tt := range goodCases {
		t.Run(tt.name, func(t *testing.T) {
			p := CartRequestPayload{tt.name, tt.Quantity}
			err := e.Register(p, ExampleCartDB{}, ctx)
			if err != nil {
				t.Errorf("error")
			}
		})
	}
	badCases := []struct {
		name     string
		ItemId   string
		Quantity int
	}{
		{
			"0はエラー", "b", 0,
		},
		{
			"負数なのでエラー", "c", -1,
		},
	}
	for _, tt := range badCases {
		t.Run(tt.name, func(t *testing.T) {
			p := CartRequestPayload{tt.name, tt.Quantity}
			err := e.Register(p, ExampleCartDB{}, ctx)
			if err == nil {
				t.Errorf("error")
			}
		})
	}
}

func TestCartUtils_InspectCart(t *testing.T) {
	utils := new(CartUtils)
	goodCases := []struct {
		name string
		cart []internalCart
		want map[string]internalCart
	}{
		{
			name: "正常",
			cart: []internalCart{{
				Cart: Cart{
					ItemId:   "1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: CartItemStatusAvailable,
						},
					},
				},
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			}},
			want: map[string]internalCart{
				"1": {
					Cart: Cart{
						ItemId:   "1",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					itemStock: 4,
					status:    items.ItemStatusAvailable,
				}},
		},
		{
			name: "在庫足りない",
			cart: []internalCart{{
				Cart: Cart{
					ItemId:   "1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: CartItemStatusAvailable,
						},
					},
				},
				itemStock: 1,
				status:    items.ItemStatusAvailable,
			}},
			want: map[string]internalCart{
				"1": {
					Cart: Cart{
						ItemId:   "1",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusStockOver,
							},
						},
					},
					itemStock: 1,
					status:    items.ItemStatusAvailable,
				}},
		},
		{
			name: "在庫なし",
			cart: []internalCart{{
				Cart: Cart{
					ItemId:   "1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: CartItemStatusAvailable,
						},
					},
				},
				itemStock: 0,
				status:    items.ItemStatusAvailable,
			}},
			want: map[string]internalCart{
				"1": {
					Cart: Cart{
						ItemId:   "1",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusNoStock,
							},
						},
					},
					itemStock: 0,
					status:    items.ItemStatusAvailable,
				}},
		},
		{
			name: "無効な商品",
			cart: []internalCart{{
				Cart: Cart{
					ItemId:   "1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: CartItemStatusAvailable,
						},
					},
				},
				itemStock: 4,
				status:    items.ItemStatusExpired,
			}},
			want: map[string]internalCart{
				"1": {
					Cart: Cart{
						ItemId:   "1",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusInvalidItem,
							},
						},
					},
					itemStock: 4,
					status:    items.ItemStatusExpired,
				}},
		},
		{
			name: "無効な商品で在庫なしの場合は無効な商品がエラーとして優先される",
			cart: []internalCart{{
				Cart: Cart{
					ItemId:   "1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: CartItemStatusAvailable,
						},
					},
				},
				itemStock: 0,
				status:    items.ItemStatusExpired,
			}},
			want: map[string]internalCart{
				"1": {
					Cart: Cart{
						ItemId:   "1",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusInvalidItem,
							},
						},
					},
					itemStock: 0,
					status:    items.ItemStatusExpired,
				}},
		},
		{
			name: "同じ商品が2つ登録されている場合に一つとして表示されるか",
			cart: []internalCart{{
				Cart: Cart{
					ItemId:   "1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: CartItemStatusAvailable,
						},
					},
				},
				itemStock: 7, //上書きされる
				status:    items.ItemStatusAvailable,
			}, {
				Cart: Cart{
					ItemId:   "1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: CartItemStatusAvailable,
						},
					},
				},
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			}},
			want: map[string]internalCart{
				"1": {
					Cart: Cart{
						ItemId:   "1",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					itemStock: 4,
					status:    items.ItemStatusAvailable,
				}},
		},
	}
	for _, tt := range goodCases {
		t.Run(tt.name, func(t *testing.T) {
			internalCart := tt.cart
			inspectedCart, _ := utils.InspectCart(internalCart)

			for internalCart := range tt.want {
				if inspectedCart[internalCart] != tt.want[internalCart] {
					t.Errorf("error")

				}
			}

		})
	}

}
