package cart

import (
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func TestCartRequests(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserDB := users.UserDB{DB: db}
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

	CartRequests := new(Requests)
	CartUtils := new(Utils)

	Cases := []struct {
		name    string
		payload []CartRequestPayload
		want    *[]Cart
		err     error
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
			want: &[]Cart{
				{
					ItemId:   "test1",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test1",
						Price: 2000,
					},
				},
				{
					ItemId:   "test2",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test2",
						Price: 3000,
					},
				},
			},
		},
		{
			name: "正常",
			payload: []CartRequestPayload{
				{
					ItemId:   "test1",
					Quantity: 2,
				},
				{
					ItemId:   "test1",
					Quantity: 1,
				},
			},
			want: &[]Cart{
				{
					ItemId:   "test1",
					Quantity: 1,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test1",
						Price: 2000,
					},
				},
			},
		},
	}

	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			for _, p := range tt.payload {
				CartRequests.Register("aaa", p, CartDB, CartUtils)
			}
			result, err := CartRequests.Get("aaa", CartDB, CartUtils)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, result, tt.want)
			}
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("%v,got,%v,want%v", tt.name, err, tt.err)
				}
			}
			for _, p := range tt.payload {
				CartRequests.Delete("aaa", p.ItemId, CartDB, CartUtils)
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
