package cash

import (
	"log"
	"net/http/httptest"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// func (StripeRequests StripeRequests) GetClientSecret(ctx *gin.Context, CartRequests cart.ICartRequests, CartDB cart.ICartDB, CartUtils cart.ICartUtils) (*string, error) {
func TestGetClientSecret(t *testing.T) {
	StripeRequests := new(StripeRequests)
	CartRequests := new(cart.CartRequests)
	CartDB := new(cart.ExampleCartDB)
	CartUtils := new(cart.CartUtils)
	Cases := []struct {
		name        string
		cart        *[]cart.InternalCart
		want        error
		SelectError error
	}{
		{
			name: "正常",
			cart: &[]cart.InternalCart{
				{
					Cart: cart.Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: cart.CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: cart.CartItemPreviewDetails{
								Status: cart.CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			want: nil,
		},
		{
			name:        "DBエラー",
			cart:        nil,
			want:        &utils.InternalError{Message: utils.InternalErrorDB},
			SelectError: &utils.InternalError{Message: utils.InternalErrorDB},
		},
		{
			name: "不正なカート",
			cart: &[]cart.InternalCart{
				{
					Cart: cart.Cart{
						ItemId:   "1",
						Quantity: 2,
						ItemProperties: cart.CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: cart.CartItemPreviewDetails{
								Status: cart.CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusExpired,
				},
				{
					Cart: cart.Cart{
						ItemId:   "2",
						Quantity: 2,
						ItemProperties: cart.CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: cart.CartItemPreviewDetails{
								Status: cart.CartItemStatusAvailable,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			want: &utils.InternalError{Message: utils.InternalErrorInvalidCart},
		},
		{
			name: "在庫オーバー",
			cart: &[]cart.InternalCart{
				{
					Cart: cart.Cart{
						ItemId:   "2",
						Quantity: 5,
						ItemProperties: cart.CartItemPreviewProperties{
							Name:  "test",
							Price: 2000,
							Details: cart.CartItemPreviewDetails{
								Status: cart.CartItemStatusStockOver,
							},
						},
					},
					ItemStock: 4,
					Status:    items.ItemStatusAvailable,
				},
			},
			want: &utils.InternalError{Message: utils.InternalErrorInvalidCart},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			CartDB.SelectError = tt.SelectError
			CartDB.InternalCarts = tt.cart
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			_, err := StripeRequests.GetClientSecret(ctx, CartRequests, CartDB, CartUtils, "test")

			if err != nil {
				log.Print(err)
				if utils.InternalErrorMessage(err.Error()) != utils.InternalErrorMessage(tt.want.Error()) {
					t.Errorf("%v,got,%v,want%v", tt.name, err.Error(), tt.want)
				}
			}

		})
	}
}
