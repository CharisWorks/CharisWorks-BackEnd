package manufacturer

import (
	"log"
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func Test_ManufacturerDB(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserDB := users.UserDB{DB: db}
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
	if err = UserDB.CreateUser("aaa"); err != nil {
		t.Errorf("error")
	}
	if err = UserDB.UpdateProfile("aaa", map[string]interface{}{"stripe_account_id": "acct_abcd"}); err != nil {
		t.Errorf("error")
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err = ManufacturerDB.RegisterItem("test", tt.payload, "aaa")
			if err != nil {
				log.Print("error", err.Error())
				t.Errorf("error")
			}
			ItemOverview, err := ItemDB.GetItemOverview("test")
			if err != nil {
				t.Errorf("error")
			}
			if !reflect.DeepEqual(*ItemOverview, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *ItemOverview, tt.want)
			}
			err = ManufacturerDB.UpdateItem(tt.updatePayload, "test")
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
	UserDB := users.UserDB{DB: db}
	ManufacturerDB := ManufacturerDB{DB: db}
	ItemDB := items.ItemDB{DB: db}
	Items := []ItemRegisterPayload{
		{
			Name:  "test1",
			Price: 2000,
			Details: ItemRegisterDetailsPayload{
				Stock:       2,
				Size:        3,
				Description: "test",
				Tags:        []string{"aaa", "bbb"},
			},
		},
		{
			Name:  "test2",
			Price: 3000,
			Details: ItemRegisterDetailsPayload{
				Stock:       3,
				Size:        4,
				Description: "test",
				Tags:        []string{"aaa", "ccc"},
			},
		},
		{
			Name:  "test3",
			Price: 4000,
			Details: ItemRegisterDetailsPayload{
				Stock:       4,
				Size:        5,
				Description: "test",
				Tags:        []string{"aaa", "ddd"},
			},
		},
		{
			Name:  "test4",
			Price: 4000,
			Details: ItemRegisterDetailsPayload{
				Stock:       4,
				Size:        5,
				Description: "test",
				Tags:        []string{"eee", "ddd"},
			},
		},
		{
			Name:  "test5",
			Price: 4000,
			Details: ItemRegisterDetailsPayload{
				Stock:       4,
				Size:        5,
				Description: "test",
				Tags:        []string{"fff", "ddd"},
			},
		},
		{
			Name:  "test6",
			Price: 5000,
			Details: ItemRegisterDetailsPayload{
				Stock:       4,
				Size:        5,
				Description: "test",
				Tags:        []string{"ggg", "ddd"},
			},
		},
	}

	if err = UserDB.CreateUser("aaa"); err != nil {
		t.Errorf("error")
	}
	for _, item := range Items {
		err = ManufacturerDB.RegisterItem(item.Name, item, "aaa")
		if err != nil {
			t.Errorf("error")
		}
	}
	Cases := []struct {
		name          string
		pageNum       int
		pageSize      int
		condition     map[string]interface{}
		tags          []string
		want          []items.ItemPreview
		totalElements int
	}{
		{
			name:      "タグのみで絞り込み",
			pageNum:   1,
			pageSize:  5,
			condition: map[string]interface{}{},
			tags:      []string{"ddd"},
			want: []items.ItemPreview{
				{
					Item_id: "test3",
					Properties: items.ItemPreviewProperties{
						Name:  "test3",
						Price: 4000,
						Details: items.ItemPreviewDetails{
							Status: items.ItemStatusReady,
						},
					},
				},
				{
					Item_id: "test4",
					Properties: items.ItemPreviewProperties{
						Name:  "test4",
						Price: 4000,
						Details: items.ItemPreviewDetails{
							Status: items.ItemStatusReady,
						},
					},
				}, {
					Item_id: "test5",
					Properties: items.ItemPreviewProperties{
						Name:  "test5",
						Price: 4000,
						Details: items.ItemPreviewDetails{
							Status: items.ItemStatusReady,
						},
					},
				}, {
					Item_id: "test6",
					Properties: items.ItemPreviewProperties{
						Name:  "test6",
						Price: 5000,
						Details: items.ItemPreviewDetails{
							Status: items.ItemStatusReady,
						},
					},
				},
			},
			totalElements: 4,
		},
		{
			name:      "条件のみで絞り込み",
			pageNum:   1,
			pageSize:  5,
			condition: map[string]interface{}{"price > ?": 4000},
			tags:      []string{},
			want: []items.ItemPreview{
				{
					Item_id: "test6",
					Properties: items.ItemPreviewProperties{
						Name:  "test6",
						Price: 5000,
						Details: items.ItemPreviewDetails{
							Status: items.ItemStatusReady,
						},
					},
				},
			},
			totalElements: 1,
		},
		{
			name:      "条件とタグで絞り込み",
			pageNum:   1,
			pageSize:  5,
			condition: map[string]interface{}{"price > ?": 3000},
			tags:      []string{"eee"},
			want: []items.ItemPreview{
				{
					Item_id: "test4",
					Properties: items.ItemPreviewProperties{
						Name:  "test4",
						Price: 4000,
						Details: items.ItemPreviewDetails{
							Status: items.ItemStatusReady,
						},
					},
				},
			},
			totalElements: 1,
		},
		{
			name:      "検索結果なし",
			pageNum:   1,
			pageSize:  5,
			condition: map[string]interface{}{"price > ?": 6000},
			tags:      []string{},
		},
		{
			name:     "ページング",
			pageNum:  2,
			pageSize: 2,
			want: []items.ItemPreview{
				{
					Item_id: "test3",
					Properties: items.ItemPreviewProperties{
						Name:  "test3",
						Price: 4000,
						Details: items.ItemPreviewDetails{
							Status: items.ItemStatusReady,
						},
					},
				},
				{
					Item_id: "test4",
					Properties: items.ItemPreviewProperties{
						Name:  "test4",
						Price: 4000,
						Details: items.ItemPreviewDetails{
							Status: items.ItemStatusReady,
						},
					},
				},
			},
			totalElements: 6,
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			previews, totalElements, err := ItemDB.GetPreviewList(tt.pageNum, tt.pageSize, tt.condition, tt.tags)
			log.Print("totalElements: ", totalElements)
			log.Print("pre: ", *previews)
			if err != nil {
				t.Errorf("error")
			}
			if !reflect.DeepEqual(*previews, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, *previews, tt.want)
			}
			if totalElements != tt.totalElements {
				t.Errorf("%v,got,%v,want%v", tt.name, totalElements, tt.totalElements)
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
