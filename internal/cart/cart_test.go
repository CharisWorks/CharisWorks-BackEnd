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
	UserDB := users.UserRepository{DB: db}
	ManufacturerDB := manufacturer.Repository{DB: db}
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

	if err = UserDB.CreateUser("aaa"); err != nil {
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

	CartRequests := Requests{CartRepository: Repository{DB: db}, CartUtils: Utils{}}

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
				CartRequests.Register("aaa", p)
			}
			result, err := CartRequests.Get("aaa")
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, result, tt.want)
			}
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("%v,got,%v,want%v", tt.name, err, tt.err)
				}
			}
			for _, p := range tt.payload {
				CartRequests.Delete("aaa", p.ItemId)
			}

		})
	}
	for _, item := range Items {
		err = ManufacturerDB.Delete(item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}
	err = UserDB.DeleteUser("aaa")
	if err != nil {
		t.Errorf("error")
	}

}
