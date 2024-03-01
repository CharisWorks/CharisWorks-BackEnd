package cart

import (
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func Test_CartCRUD(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserDB := user.UserDB{DB: db}
	ManufacturerDB := manufacturer.ManufacturerDB{DB: db}
	CartDB := CartDB{DB: db}
	Items := []manufacturer.ItemRegisterPayload{
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
		{
			Name:  "test3",
			Price: 4000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       4,
				Size:        5,
				Description: "test",
				Tags:        []string{"aaa", "ddd"},
			},
		},
		{
			Name:  "test4",
			Price: 4000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       4,
				Size:        5,
				Description: "test",
				Tags:        []string{"eee", "ddd"},
			},
		},
		{
			Name:  "test5",
			Price: 4000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       4,
				Size:        5,
				Description: "test",
				Tags:        []string{"fff", "ddd"},
			},
		},
		{
			Name:  "test6",
			Price: 5000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       4,
				Size:        5,
				Description: "test",
				Tags:        []string{"ggg", "ddd"},
			},
		},
	}

	if err = UserDB.CreateUser("aaa", 1); err != nil {
		t.Errorf("error")
	}
	for _, item := range Items {
		err = ManufacturerDB.RegisterItem(item.Name, item, 1, "aaa")
		if err != nil {
			t.Errorf("error")
		}
		err = ManufacturerDB.UpdateItem(map[string]interface{}{"status": items.ItemStatusAvailable}, 2, item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}
	Cases := []struct {
		name          string
		payload       []CartRequestPayload
		want          []InternalCart
		updatePayload CartRequestPayload
		wantUpdated   []InternalCart
	}{
		{
			name: "正常",
			payload: []CartRequestPayload{
				{
					ItemId:   "test1",
					Quantity: 2,
				},
				{
					ItemId:   "test2",
					Quantity: 2,
				},
			},
			want: []InternalCart{
				{
					Index: 0,
					Cart: Cart{
						ItemId:   "test1",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test1",
							Price: 2000,
						},
					},

					ItemStock: 2,
					Status:    items.ItemStatusAvailable,
				},
				{
					Index: 1,
					Cart: Cart{
						ItemId:   "test2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test2",
							Price: 3000,
						},
					},
					ItemStock: 3,
					Status:    items.ItemStatusAvailable,
				},
			},
			updatePayload: CartRequestPayload{
				ItemId:   "test1",
				Quantity: 3,
			},
			wantUpdated: []InternalCart{
				{
					Index: 0,
					Cart: Cart{
						ItemId:   "test1",
						Quantity: 3,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test1",
							Price: 2000,
						},
					},

					ItemStock: 2,
					Status:    items.ItemStatusAvailable,
				},
				{
					Index: 1,
					Cart: Cart{
						ItemId:   "test2",
						Quantity: 2,
						ItemProperties: CartItemPreviewProperties{
							Name:  "test2",
							Price: 3000,
						},
					},
					ItemStock: 3,
					Status:    items.ItemStatusAvailable,
				},
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			for _, p := range tt.payload {
				err := CartDB.RegisterCart("aaa", p)
				if err != nil {
					t.Errorf(err.Error())
				}
			}
			Cart, err := CartDB.GetCart("aaa")
			if err != nil {
				t.Errorf(err.Error())
			}
			if !reflect.DeepEqual(*Cart, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *Cart, tt.want)
			}

			err = CartDB.UpdateCart("aaa", tt.updatePayload)
			if err != nil {
				t.Errorf(err.Error())
			}
			Cart, err = CartDB.GetCart("aaa")
			if err != nil {
				t.Errorf(err.Error())
			}
			if !reflect.DeepEqual(*Cart, tt.wantUpdated) {
				t.Errorf("%v,got,%v,want%v", tt.name, *Cart, tt.wantUpdated)
			}
			for _, p := range tt.payload {
				err := CartDB.DeleteCart("aaa", p.ItemId)
				if err != nil {
					t.Errorf(err.Error())
				}
			}

		})
	}
	for _, item := range Items {
		err = ManufacturerDB.DeleteItem(item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}
	err = UserDB.DeleteUser("aaa")
	if err != nil {
		t.Errorf("error")
	}
}
