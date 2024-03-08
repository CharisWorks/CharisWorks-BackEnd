package transaction

import (
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func Test_CartCRUD(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}

	UserDB := users.UserRepository{DB: db}
	ManufacturerDB := manufacturer.Repository{DB: db}
	cartRepository := cart.Repository{DB: db}
	Items := []manufacturer.RegisterPayload{
		{
			Name:  "test1",
			Price: 2000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       2,
				Size:        3,
				Description: "test",
				Tags:        []string{"aaa", "bbb"},
			},
		},
		{
			Name:  "test2",
			Price: 3000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       3,
				Size:        4,
				Description: "test",
				Tags:        []string{"aaa", "ccc"},
			},
		},
	}
	if err = UserDB.Create("aaa"); err != nil {
		t.Errorf("error")
	}
	if err = UserDB.UpdateProfile("aaa", map[string]interface{}{
		"display_name":      "test",
		"description":       "test",
		"stripe_account_id": "test",
	}); err != nil {
		t.Errorf("error")
	}
	for _, item := range Items {
		err = ManufacturerDB.Register(item.Name, item, "aaa")
		if err != nil {
			t.Errorf("error")
		}
		err = ManufacturerDB.Update(map[string]interface{}{"status": items.Available}, item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}
	Cases := []struct {
		name          string
		payload       []cart.CartRequestPayload
		want          []cart.InternalCart
		updatePayload cart.CartRequestPayload
		wantUpdated   []cart.InternalCart
	}{
		{
			name: "正常",
			payload: []cart.CartRequestPayload{
				{
					ItemId:   "test1",
					Quantity: 2,
				},
				{
					ItemId:   "test2",
					Quantity: 2,
				},
			},
			want: []cart.InternalCart{
				{
					Index: 0,
					Cart: cart.Cart{
						ItemId:   "test1",
						Quantity: 2,
						ItemProperties: cart.CartItemPreviewProperties{
							Name:  "test1",
							Price: 2000,
						},
					},
					Item: cart.InternalItem{
						Price:                   2000,
						Name:                    "test1",
						Description:             "test",
						Tags:                    []string{"aaa", "bbb"},
						Size:                    3,
						ManufacturerUserId:      "aaa",
						ManufacturerName:        "test",
						ManufacturerDescription: "test",
					},
					ItemStock: 2,
					Status:    items.Available,
				},
				{
					Index: 1,
					Cart: cart.Cart{
						ItemId:   "test2",
						Quantity: 2,
						ItemProperties: cart.CartItemPreviewProperties{
							Name:  "test2",
							Price: 3000,
						},
					},
					Item: cart.InternalItem{
						Price:                   2000,
						Name:                    "test1",
						Description:             "test",
						Tags:                    []string{"aaa", "bbb"},
						Size:                    3,
						ManufacturerUserId:      "aaa",
						ManufacturerName:        "test",
						ManufacturerDescription: "test",
					},
					ItemStock: 3,
					Status:    items.Available,
				},
			},
			updatePayload: cart.CartRequestPayload{
				ItemId:   "test1",
				Quantity: 3,
			},
			wantUpdated: []cart.InternalCart{
				{
					Index: 0,
					Cart: cart.Cart{
						ItemId:   "test1",
						Quantity: 3,
						ItemProperties: cart.CartItemPreviewProperties{
							Name:  "test1",
							Price: 2000,
						},
					},
					Item: cart.InternalItem{
						Price:                   2000,
						Name:                    "test1",
						Description:             "test",
						Tags:                    []string{"aaa", "bbb"},
						Size:                    3,
						ManufacturerUserId:      "aaa",
						ManufacturerName:        "test",
						ManufacturerDescription: "test",
					},
					ItemStock: 2,
					Status:    items.Available,
				},
				{
					Index: 1,
					Cart: cart.Cart{
						ItemId:   "test2",
						Quantity: 2,
						ItemProperties: cart.CartItemPreviewProperties{
							Name:  "test2",
							Price: 3000,
						},
					},
					Item: cart.InternalItem{
						Price:                   3000,
						Name:                    "test2",
						Description:             "test",
						Tags:                    []string{"aaa", "ccc"},
						Size:                    4,
						ManufacturerUserId:      "aaa",
						ManufacturerName:        "test",
						ManufacturerDescription: "test",
					},
					ItemStock: 3,
					Status:    items.Available,
				},
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			for _, p := range tt.payload {
				err := cartRepository.Register("aaa", p)
				if err != nil {
					t.Errorf(err.Error())
				}
			}
			Cart, err := cartRepository.Get("aaa")
			if err != nil {
				t.Errorf(err.Error())
			}
			if Cart == &tt.want {
				t.Errorf("%v,got,%v,want%v", tt.name, *Cart, tt.want)
			}

			err = cartRepository.Update("aaa", tt.updatePayload)
			if err != nil {
				t.Errorf(err.Error())
			}
			Cart, err = cartRepository.Get("aaa")
			if err != nil {
				t.Errorf(err.Error())
			}
			if !reflect.DeepEqual(*Cart, tt.wantUpdated) {
				t.Errorf("%v,got,%v,want%v", tt.name, *Cart, tt.wantUpdated)
			}
			for _, p := range tt.payload {
				err := cartRepository.Delete("aaa", p.ItemId)
				if err != nil {
					t.Errorf(err.Error())
				}
			}

		})
	}
	for _, item := range Items {
		err = ManufacturerDB.Delete(item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}
	err = UserDB.Delete("aaa")
	if err != nil {
		t.Errorf("error")
	}
}
