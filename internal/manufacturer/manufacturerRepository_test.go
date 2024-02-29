package manufacturer

import (
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func Test_ManufacturerDB(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}

	UserDB := user.UserDB{DB: db}
	ManufacturerDB := ManufacturerDB{DB: db}
	ItemDB := items.ItemDB{DB: db}
	Cases := []struct {
		name          string
		payload       ItemRegisterPayload
		want          items.ItemOverview
		updatePayload map[string]interface{}
		wantUpdated   items.ItemOverview
	}{
		{
			name: "正常",
			payload: ItemRegisterPayload{
				Name:  "abc",
				Price: 2000,
				Details: ItemRegisterDetailsPayload{
					Stock:       2,
					Size:        3,
					Description: "test",
					Tags:        []string{"aaa", "bbb"},
				},
			},
			want: items.ItemOverview{
				Item_id: "aaa",
				Properties: items.ItemOverviewProperties{
					Name:  "abc",
					Price: 2000,
					Details: items.ItemOverviewDetails{
						Status:      items.ItemStatusReady,
						Stock:       2,
						Size:        3,
						Description: "test",
						Tags:        []string{"aaa", "bbb"},
					},
				},
			},
			updatePayload: map[string]interface{}{
				"stock": 4,
			},
			wantUpdated: items.ItemOverview{
				Item_id: "aaa",
				Properties: items.ItemOverviewProperties{
					Name:  "abc",
					Price: 2000,
					Details: items.ItemOverviewDetails{
						Status:      items.ItemStatusReady,
						Stock:       4,
						Size:        3,
						Description: "test",
						Tags:        []string{"aaa", "bbb"},
					},
				},
			},
		},
	}
	if err = UserDB.CreateUser("aaa", 1); err != nil {
		t.Errorf("error")
	}
	if err = UserDB.UpdateProfile("aaa", map[string]interface{}{"stripe_account_id": "acct_abcd"}); err != nil {
		t.Errorf("error")
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err = ManufacturerDB.RegisterItem(tt.payload, 1, "aaa")
			if err != nil {
				t.Errorf("error")
			}
			ItemOverview, err := ItemDB.GetItemOverview(1)
			if err != nil {
				t.Errorf("error")
			}

			if !reflect.DeepEqual(*ItemOverview, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *ItemOverview, tt.want)
			}
			err = ManufacturerDB.UpdateItem(tt.updatePayload, 1, 1)
			if err != nil {
				t.Errorf("error")
			}
			ItemOverview, err = ItemDB.GetItemOverview(1)
			if err != nil {
				t.Errorf("error")
			}
			if !reflect.DeepEqual(*ItemOverview, tt.wantUpdated) {
				t.Errorf("%v,got,%v,want%v", tt.name, *ItemOverview, tt.wantUpdated)
			}
			err = ManufacturerDB.DeleteItem(1)
			if err != nil {
				t.Errorf("error")
			}

		})
	}
	err = UserDB.DeleteUser("aaa")
	if err != nil {
		t.Errorf("error")
	}
}
