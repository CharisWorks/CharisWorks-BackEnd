package cart

import (
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func TestCartUtils_InspectCart(t *testing.T) {
	CartUtils := new(Utils)
	Cases := []struct {
		name string
		cart []InternalCart
		want map[string]InternalCart
		err  utils.InternalMessage
	}{
		{
			name: "正常",
			cart: []InternalCart{{
				Cart: Cart{
					ItemId:   "test",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
				ItemStock: 4,
				Status:    items.Available,
			}},
			want: map[string]InternalCart{
				"test": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: Available,
							},
						},
					},
					ItemStock: 4,
					Status:    items.Available,
				}},
		},
		{
			name: "在庫足りない",
			cart: []InternalCart{{
				Cart: Cart{
					ItemId:   "test",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
				ItemStock: 1,
				Status:    items.Available,
			}},
			want: map[string]InternalCart{
				"test": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: StockOver,
							},
						},
					},
					ItemStock: 1,
					Status:    items.Available,
				}},
			err: utils.InternalErrorInvalidCart,
		},
		{
			name: "在庫なし",
			cart: []InternalCart{{
				Cart: Cart{
					ItemId:   "test",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
				ItemStock: 0,
				Status:    items.Available,
			}},
			want: map[string]InternalCart{
				"test": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: NoStock,
							},
						},
					},
					ItemStock: 0,
					Status:    items.Available,
				}},
			err: utils.InternalErrorInvalidCart,
		},
		{
			name: "無効な商品",
			cart: []InternalCart{{
				Cart: Cart{
					ItemId:   "test",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
				ItemStock: 4,
				Status:    items.Expired,
			}},
			want: map[string]InternalCart{
				"test": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: InvalidItem,
							},
						},
					},
					ItemStock: 4,
					Status:    items.Expired,
				}},
			err: utils.InternalErrorInvalidCart,
		},
		{
			name: "無効な商品で在庫なしの場合は無効な商品がエラーとして優先される",
			cart: []InternalCart{{
				Cart: Cart{
					ItemId:   "test",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
				ItemStock: 0,
				Status:    items.Expired,
			}},
			want: map[string]InternalCart{
				"test": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: InvalidItem,
							},
						},
					},
					ItemStock: 0,
					Status:    items.Expired,
				}},
			err: utils.InternalErrorInvalidCart,
		},
		{
			name: "同じ商品が2つ登録されている場合に一つとして表示されるか",
			cart: []InternalCart{{
				Cart: Cart{
					ItemId:   "test",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
				ItemStock: 7, //上書きされる
				Status:    items.Available,
			}, {
				Cart: Cart{
					ItemId:   "test",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
				ItemStock: 4,
				Status:    items.Available,
			}},
			want: map[string]InternalCart{
				"test": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: Available,
							},
						},
					},
					ItemStock: 4,
					Status:    items.Available,
				}},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			internalCart := tt.cart
			inspectedCart, err := CartUtils.Inspect(internalCart)
			for internalCart := range tt.want {
				if inspectedCart[internalCart] != tt.want[internalCart] {
					t.Errorf("%v,got,%v,want%v", tt.name, inspectedCart[internalCart], tt.want[internalCart])

				}
			}
			if err != nil {
				if utils.InternalMessage(err.Error()) != tt.err {
					t.Errorf("%v,got,%v,want%v", tt.name, err.Error(), tt.err)
				}
			}

		})
	}

}

func TestCartUtils_InspectPayload(t *testing.T) {
	e := new(Utils)
	Cases := []struct {
		name    string
		Payload CartRequestPayload
		Status  items.ItemStatus
		want    *CartRequestPayload
		err     utils.InternalMessage
	}{
		{
			name: "正常なパターン",
			Payload: CartRequestPayload{
				ItemId:   "test",
				Quantity: 2,
			},
			Status: items.ItemStatus{
				Stock:  3,
				Status: items.Available,
			},
			want: &CartRequestPayload{
				ItemId:   "test",
				Quantity: 2,
			},
		}, {
			name: "在庫足りない",
			Payload: CartRequestPayload{
				ItemId:   "test",
				Quantity: 2,
			},
			Status: items.ItemStatus{
				Stock:  1,
				Status: items.Available,
			},
			want: nil,
			err:  utils.InternalErrorStockOver,
		}, {
			name: "在庫ない",
			Payload: CartRequestPayload{
				ItemId:   "test",
				Quantity: 2,
			},
			Status: items.ItemStatus{
				Stock:  0,
				Status: items.Available,
			},
			want: nil,
			err:  utils.InternalErrorNoStock,
		}, {
			name: "無効な商品",
			Payload: CartRequestPayload{
				ItemId:   "test",
				Quantity: 2,
			},
			Status: items.ItemStatus{
				Stock:  4,
				Status: items.Expired,
			},
			want: nil,
			err:  utils.InternalErrorInvalidItem,
		}, {
			name: "在庫切れだけど無効な商品だと無効な商品のエラーを出す",
			Payload: CartRequestPayload{
				ItemId:   "test",
				Quantity: 2,
			},
			Status: items.ItemStatus{
				Stock:  0,
				Status: items.Expired,
			},
			want: nil,
			err:  utils.InternalErrorInvalidItem,
		}, {
			name: "無効なペイロード(0)",
			Payload: CartRequestPayload{
				ItemId:   "test",
				Quantity: 0,
			},
			Status: items.ItemStatus{
				Stock:  4,
				Status: items.Expired,
			},
			want: nil,
			err:  utils.InternalErrorInvalidPayload,
		}, {
			name: "無効なペイロード(負数)",
			Payload: CartRequestPayload{
				ItemId:   "test",
				Quantity: -3,
			},
			Status: items.ItemStatus{
				Stock:  4,
				Status: items.Expired,
			},
			want: nil,
			err:  utils.InternalErrorInvalidPayload,
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			InspectedPayload, err := e.InspectPayload(tt.Payload, tt.Status)
			if !reflect.DeepEqual(InspectedPayload, tt.want) {
				t.Errorf("want %v, got %v", tt.want, InspectedPayload)
			}
			if err != nil {
				if utils.InternalMessage(err.Error()) != tt.err {
					t.Errorf("want %v, got %v", tt.err, err.Error())
				}
			}
		})
	}
}

func TestCartUtils_ConvertCart(t *testing.T) {
	CartUtils := new(Utils)
	Cases := []struct {
		name          string
		inspectedCart map[string]InternalCart
		want          []Cart
	}{
		{
			name: "正常",
			inspectedCart: map[string]InternalCart{
				"1": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: Available,
							},
						},
					},
					ItemStock: 4,
					Status:    items.Available,
				},
			},
			want: []Cart{
				{
					ItemId:   "test",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test",
						Price: 2000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			Cart := CartUtils.Convert(tt.inspectedCart)
			if !reflect.DeepEqual(Cart, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, Cart, tt.want)

			}

		})
	}

}
func TestCartUtils_GetTotalAmount(t *testing.T) {
	CartUtils := new(Utils)
	Cases := []struct {
		name          string
		inspectedCart map[string]InternalCart
		want          int
	}{
		{
			name: "1つパターン",
			inspectedCart: map[string]InternalCart{
				"1": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: Available,
							},
						},
					},
					ItemStock: 4,
					Status:    items.Available,
				},
			},
			want: 4000,
		},
		{
			name: "2つパターン",
			inspectedCart: map[string]InternalCart{
				"1": {
					Cart: Cart{
						ItemId:   "test",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: Available,
							},
						},
					},
					ItemStock: 4,
					Status:    items.Available,
				},
				"2": {
					Cart: Cart{
						ItemId:   "test2",
						Quantity: 1,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 5000,
							Details: CartItemPreviewDetails{
								Status: Available,
							},
						},
					},
					ItemStock: 4,
					Status:    items.Available,
				},
			},
			want: 9000,
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			totalAmount := CartUtils.GetTotalAmount(tt.inspectedCart)
			if totalAmount != tt.want {
				t.Errorf("%v,got,%v,want%v", tt.name, totalAmount, tt.want)

			}

		})
	}

}
