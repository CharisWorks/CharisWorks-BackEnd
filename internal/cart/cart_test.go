package cart

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func TestCartUtils_InspectCart(t *testing.T) {
	CartUtils := new(CartUtils)
	Cases := []struct {
		name string
		cart []InternalCart
		want map[string]InternalCart
		err  utils.InternalErrorMessage
	}{
		{
			name: "正常",
			cart: []InternalCart{{
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
				ItemStock: 4,
				Status:    items.ItemStatusAvailable,
			}},
			want: map[string]InternalCart{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				}},
		},
		{
			name: "在庫足りない",
			cart: []InternalCart{{
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
				ItemStock: 1,
				Status:    items.ItemStatusAvailable,
			}},
			want: map[string]InternalCart{
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
					ItemStock: 1,
					Status:    items.ItemStatusAvailable,
				}},
			err: utils.InternalErrorInvalidCart,
		},
		{
			name: "在庫なし",
			cart: []InternalCart{{
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
				ItemStock: 0,
				Status:    items.ItemStatusAvailable,
			}},
			want: map[string]InternalCart{
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
					ItemStock: 0,
					Status:    items.ItemStatusAvailable,
				}},
			err: utils.InternalErrorInvalidCart,
		},
		{
			name: "無効な商品",
			cart: []InternalCart{{
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
				ItemStock: 4,
				Status:    items.ItemStatusExpired,
			}},
			want: map[string]InternalCart{
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
					ItemStock: 4,
					Status:    items.ItemStatusExpired,
				}},
			err: utils.InternalErrorInvalidCart,
		},
		{
			name: "無効な商品で在庫なしの場合は無効な商品がエラーとして優先される",
			cart: []InternalCart{{
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
				ItemStock: 0,
				Status:    items.ItemStatusExpired,
			}},
			want: map[string]InternalCart{
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
					ItemStock: 0,
					Status:    items.ItemStatusExpired,
				}},
			err: utils.InternalErrorInvalidCart,
		},
		{
			name: "同じ商品が2つ登録されている場合に一つとして表示されるか",
			cart: []InternalCart{{
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
				ItemStock: 7, //上書きされる
				Status:    items.ItemStatusAvailable,
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
				ItemStock: 4,
				Status:    items.ItemStatusAvailable,
			}},
			want: map[string]InternalCart{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				}},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			internalCart := tt.cart
			inspectedCart, err := CartUtils.InspectCart(internalCart)
			for internalCart := range tt.want {
				if inspectedCart[internalCart] != tt.want[internalCart] {
					t.Errorf("%v,got,%v,want%v", tt.name, inspectedCart[internalCart], tt.want[internalCart])

				}
			}
			if err != nil {
				if utils.InternalErrorMessage(err.Error()) != tt.err {
					t.Errorf("%v,got,%v,want%v", tt.name, err.Error(), tt.err)
				}
			}

		})
	}

}

func TestCartUtils_InspectPayload(t *testing.T) {
	e := new(CartUtils)
	Cases := []struct {
		name    string
		Payload CartRequestPayload
		Status  itemStatus
		want    *CartRequestPayload
		err     utils.InternalErrorMessage
	}{
		{
			name: "正常なパターン",
			Payload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			Status: itemStatus{
				itemStock: 3,
				status:    items.ItemStatusAvailable,
			},
			want: &CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
		}, {
			name: "在庫足りない",
			Payload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			Status: itemStatus{
				itemStock: 1,
				status:    items.ItemStatusAvailable,
			},
			want: nil,
			err:  utils.InternalErrorStockOver,
		}, {
			name: "在庫ない",
			Payload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			Status: itemStatus{
				itemStock: 0,
				status:    items.ItemStatusAvailable,
			},
			want: nil,
			err:  utils.InternalErrorNoStock,
		}, {
			name: "無効な商品",
			Payload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			Status: itemStatus{
				itemStock: 4,
				status:    items.ItemStatusExpired,
			},
			want: nil,
			err:  utils.InternalErrorInvalidItem,
		}, {
			name: "在庫切れだけど無効な商品だと無効な商品のエラーを出す",
			Payload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			Status: itemStatus{
				itemStock: 0,
				status:    items.ItemStatusExpired,
			},
			want: nil,
			err:  utils.InternalErrorInvalidItem,
		}, {
			name: "無効なペイロード(0)",
			Payload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 0,
			},
			Status: itemStatus{
				itemStock: 4,
				status:    items.ItemStatusExpired,
			},
			want: nil,
			err:  utils.InternalErrorInvalidPayload,
		}, {
			name: "無効なペイロード(負数)",
			Payload: CartRequestPayload{
				ItemId:   "1",
				Quantity: -3,
			},
			Status: itemStatus{
				itemStock: 4,
				status:    items.ItemStatusExpired,
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
				if utils.InternalErrorMessage(err.Error()) != tt.err {
					t.Errorf("want %v, got %v", tt.err, err.Error())
				}
			}
		})
	}
}

func TestCartUtils_ConvertCart(t *testing.T) {
	CartUtils := new(CartUtils)
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			want: []Cart{
				{
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
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			Cart := CartUtils.ConvertCart(tt.inspectedCart)
			if !reflect.DeepEqual(Cart, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, Cart, tt.want)

			}

		})
	}

}
func TestCartUtils_GetTotalAmount(t *testing.T) {
	CartUtils := new(CartUtils)
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			want: 4000,
		},
		{
			name: "2つパターン",
			inspectedCart: map[string]InternalCart{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
				"2": {
					Cart: Cart{
						ItemId:   "2",
						Quantity: 1,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 5000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
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

func TestCartRequests_Get(t *testing.T) {
	CartRequests := new(CartRequests)
	CartUtils := new(CartUtils)
	CartDB := new(ExampleCartDB)

	Cases := []struct {
		name          string
		internalCarts *[]InternalCart
		want          *[]Cart
		err           error
	}{
		{
			name: "正常",
			internalCarts: &[]InternalCart{
				{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			want: &[]Cart{
				{
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
			},
		},
		{
			name: "無効な商品",
			internalCarts: &[]InternalCart{
				{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			want: &[]Cart{
				{
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
			},
			err: &utils.InternalError{Message: utils.InternalErrorInvalidCart},
		},
		{
			name:          "カートが存在しない",
			internalCarts: nil,
			want:          nil,
			err:           &utils.InternalError{Message: utils.InternalErrorNotFound},
		},
		{
			name: "無効なカート",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "1",
						Quantity: 0,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 0,
					Status:    items.ItemStatusAvailable,
				},
			},
			want: nil,
			err:  &utils.InternalError{Message: utils.InternalErrorInvalidCart},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {

			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			CartDB.InternalCarts = tt.internalCarts
			result, err := CartRequests.Get(ctx, CartDB, CartUtils, "test")
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, result, tt.want)
			}
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("%v,got,%v,want%v", tt.name, err, tt.err)
				}
			}
		})
	}

}

func TestCartRequests_Register(t *testing.T) {
	CartRequests := new(CartRequests)
	CartUtils := new(CartUtils)
	CartDB := new(ExampleCartDB)

	Cases := []struct {
		name               string
		internalCarts      *[]InternalCart
		CartRequestPayload CartRequestPayload
		itemStatus         *itemStatus
		DBerr              error
		UpdateDBerr        error
		registerDBerr      error
		err                error
	}{
		{
			name: "正常 存在する場合",
			internalCarts: &[]InternalCart{
				{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			DBerr: nil,
			err:   nil,
		},
		{
			name: "正常 存在しない場合",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			DBerr: nil,
			err:   nil,
		},
		{
			name:          "エラー カート取得失敗",
			internalCarts: nil,
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			DBerr: &utils.InternalError{Message: utils.InternalErrorNotFound},
			err:   &utils.InternalError{Message: utils.InternalErrorNotFound},
		},
		{
			name: "エラー 商品が存在しない場合",
			internalCarts: &[]InternalCart{
				{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: nil,
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			DBerr: &utils.InternalError{Message: utils.InternalErrorNotFound},
			err:   &utils.InternalError{Message: utils.InternalErrorNotFound},
		},
		{
			name: "エラー カートエラー 存在する場合",
			internalCarts: &[]InternalCart{
				{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			DBerr: nil,
			err:   &utils.InternalError{Message: utils.InternalErrorInvalidCart},
		},
		{
			name: "エラー カートエラー 存在しない場合",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusInvalidItem,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			DBerr: nil,
			err:   &utils.InternalError{Message: utils.InternalErrorInvalidCart},
		},
		{
			name: "ペイロードエラー 存在する場合",
			internalCarts: &[]InternalCart{
				{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: -1,
			},
			DBerr: nil,
			err:   &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "ペイロードエラー 存在しない場合",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 0,
			},
			DBerr: nil,
			err:   &utils.InternalError{Message: utils.InternalErrorInvalidPayload},
		},
		{
			name: "ペイロードエラー 在庫オーバーの場合",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 8,
			},
			DBerr: nil,
			err:   &utils.InternalError{Message: utils.InternalErrorStockOver},
		},
		{
			name: "update  error",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "2",
				Quantity: 2,
			},
			UpdateDBerr: &utils.InternalError{Message: utils.InternalErrorDB},
			err:         &utils.InternalError{Message: utils.InternalErrorDB},
		},
		{
			name: "register  error",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemStatus: &itemStatus{
				itemStock: 4,
				status:    items.ItemStatusAvailable,
			},
			CartRequestPayload: CartRequestPayload{
				ItemId:   "1",
				Quantity: 2,
			},
			registerDBerr: &utils.InternalError{Message: utils.InternalErrorDB},
			err:           &utils.InternalError{Message: utils.InternalErrorDB},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			CartDB.InternalCarts = tt.internalCarts
			CartDB.Err = tt.DBerr
			CartDB.UpdateErr = tt.UpdateDBerr
			CartDB.RegisterError = tt.registerDBerr
			CartDB.ItemStatus = tt.itemStatus
			err := CartRequests.Register(tt.CartRequestPayload, CartDB, CartUtils, ctx, "test")
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("%v,got,%v,want%v", tt.name, err, tt.err)
				}
			}
		})
	}

}

func TestCartRequests_Delete(t *testing.T) {
	CartRequests := new(CartRequests)
	CartUtils := new(CartUtils)
	CartDB := new(ExampleCartDB)

	Cases := []struct {
		name          string
		internalCarts *[]InternalCart
		itemId        string
		DBerr         error
		UpdateDBerr   error
		RegisterDBerr error
		DeleteDBerr   error
		err           error
	}{
		{
			name: "正常 存在する場合",
			internalCarts: &[]InternalCart{
				{
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
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemId: "1",
			DBerr:  nil,
			err:    nil,
		},
		{
			name: "エラー 対象が存在しない場合",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemId: "1",
			DBerr:  nil,
			err:    nil,
		},
		{
			name:          "エラー カート取得失敗",
			internalCarts: nil,
			itemId:        "1",
			DBerr:         &utils.InternalError{Message: utils.InternalErrorNotFound},
			err:           &utils.InternalError{Message: utils.InternalErrorNotFound},
		},
		{
			name: "delete  error",
			internalCarts: &[]InternalCart{
				{
					Cart: Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: CartItemPreviewDetails{
								Status: CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			itemId:      "2",
			DeleteDBerr: &utils.InternalError{Message: utils.InternalErrorDB},
			err:         &utils.InternalError{Message: utils.InternalErrorDB},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			CartDB.InternalCarts = tt.internalCarts
			CartDB.Err = tt.DBerr
			CartDB.DeleteError = tt.DeleteDBerr
			err := CartRequests.Delete(tt.itemId, CartDB, CartUtils, ctx, "test")

			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("%v,got,%v,want%v", tt.name, err, tt.err)
				}
			}
		})
	}

}
