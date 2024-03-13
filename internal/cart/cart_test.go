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

	CartRequests := Requests{CartRepository: Repository{DB: db}, CartUtils: Utils{}, ItemGetStatus: items.GetStatus{DB: db}}

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
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
				{
					ItemId:   "test2",
					Quantity: 2,
					ItemProperties: CartItemPreviewProperties{
						Name:  "test2",
						Price: 3000,
						Details: CartItemPreviewDetails{
							Status: Available,
						},
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
						Details: CartItemPreviewDetails{
							Status: Available,
						},
					},
				},
			},
		},
		{
			name: "在庫不足",
			payload: []CartRequestPayload{
				{
					ItemId:   "test1",
					Quantity: 3,
				},
			},
			want: nil,
			err:  &utils.InternalError{Message: utils.InternalErrorStockOver},
		},
		{
			name: "存在しない商品",
			payload: []CartRequestPayload{
				{
					ItemId:   "test3",
					Quantity: 3,
				},
			},
			want: nil,
			err:  &utils.InternalError{Message: utils.InternalErrorDB},
		},
	}

	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			for _, p := range tt.payload {
				err := CartRequests.Register("aaa", p)
				if err != nil {
					if err.Error() != tt.err.Error() {
						t.Errorf(err.Error())
					}
					return
				}
			}
			result, err := CartRequests.Get("aaa")
			if !reflect.DeepEqual(&result, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, result, tt.want)
			}
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("%v,got,%v,want%v", tt.name, err, tt.err)
				}
				return
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
	err = UserDB.Delete("aaa")
	if err != nil {
		t.Errorf("error")
	}

}
