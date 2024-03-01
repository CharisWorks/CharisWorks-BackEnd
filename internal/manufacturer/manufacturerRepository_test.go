package manufacturer

import (
	"log"
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
				Item_id: "test",
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
				Item_id: "test",
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
			err = ManufacturerDB.RegisterItem("test", tt.payload, 1, "aaa")
			if err != nil {
				t.Errorf("error")
			}
			ItemOverview, err := ItemDB.GetItemOverview("test")
			if err != nil {
				t.Errorf("error")
			}
			if !reflect.DeepEqual(*ItemOverview, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *ItemOverview, tt.want)
			}
			err = ManufacturerDB.UpdateItem(tt.updatePayload, 1, "test")
			if err != nil {
				t.Errorf("error")
			}
			ItemOverview, err = ItemDB.GetItemOverview("test")
			if err != nil {
				t.Errorf("error")
			}
			if !reflect.DeepEqual(*ItemOverview, tt.wantUpdated) {
				t.Errorf("%v,got,%v,want%v", tt.name, *ItemOverview, tt.wantUpdated)
			}
			err = ManufacturerDB.DeleteItem("test")
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
func Test_GetItemList(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserDB := user.UserDB{DB: db}
	ManufacturerDB := ManufacturerDB{DB: db}
	ItemDB := items.ItemDB{DB: db}
	Cases := []struct {
		name          string
		payload       []ItemRegisterPayload
		want          items.ItemOverview
		updatePayload map[string]interface{}
		wantUpdated   items.ItemOverview
	}{
		{
			name: "正常",
			payload: []ItemRegisterPayload{{
				Name:  "abc",
				Price: 2000,
				Details: ItemRegisterDetailsPayload{
					Stock:       2,
					Size:        3,
					Description: "test",
					Tags:        []string{"aaa", "bbb"},
				},
			},
				{
					Name:  "def",
					Price: 3000,
					Details: ItemRegisterDetailsPayload{
						Stock:       3,
						Size:        4,
						Description: "test",
						Tags:        []string{"aaa", "ccc"},
					},
				},
				{
					Name:  "ghi",
					Price: 4000,
					Details: ItemRegisterDetailsPayload{
						Stock:       4,
						Size:        5,
						Description: "test",
						Tags:        []string{"aaa", "ddd"},
					},
				},
			},
			want: items.ItemOverview{
				Item_id: "test",
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
				Item_id: "test",
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
			for _, payload := range tt.payload {
				err = ManufacturerDB.RegisterItem(payload.Name, payload, 1, "aaa")
				if err != nil {
					t.Errorf("error")
				}
			}
			ItemOverview, err := ItemDB.GetItemOverview("abc")
			if err != nil {
				t.Errorf("error")
			}
			log.Print(*ItemOverview)
			previews, err := ItemDB.GetPreviewList(1, 5, map[string]interface{}{}, []string{"ccc", "ddd"})

			log.Print("pre: ", *previews)
			if err != nil {
				t.Errorf("error")
			}

			for _, payload := range tt.payload {
				err = ManufacturerDB.DeleteItem(payload.Name)
				if err != nil {
					t.Errorf("error")
				}
			}

		})
	}
	err = UserDB.DeleteUser("aaa")
	if err != nil {
		t.Errorf("error")
	}
}
