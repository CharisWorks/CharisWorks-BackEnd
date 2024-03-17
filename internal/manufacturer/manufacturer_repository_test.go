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
	After(t)
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserDB := users.UserRepository{DB: db}
	ManufacturerDB := TestRepository{DB: db}
	ItemRepository := items.ItemRepository{DB: db}
	Cases := []struct {
		name          string
		payload       RegisterPayload
		want          items.Overview
		updatePayload map[string]interface{}
		wantUpdated   items.Overview
	}{
		{
			name: "正常",
			payload: RegisterPayload{
				Name:  "abc",
				Price: 2000,
				Details: ItemRegisterDetailsPayload{
					Stock:       2,
					Size:        3,
					Description: "test",
					Tags:        []string{"aaa", "bbb"},
				},
			},
			want: items.Overview{
				Item_id: "test",
				Properties: items.OverviewProperties{
					Name:  "abc",
					Price: 2000,
					Details: items.OverviewDetails{
						Status:      items.Ready,
						Stock:       2,
						Size:        3,
						Description: "test",
						Tags:        []string{"aaa", "bbb"},
					},
				},
				Manufacturer: items.ManufacturerDetails{
					Name:            "test",
					StripeAccountId: "test",
					Description:     "test",
					UserId:          "aaa",
				},
			},
			updatePayload: map[string]interface{}{
				"stock": 4,
			},
			wantUpdated: items.Overview{
				Item_id: "test",
				Properties: items.OverviewProperties{
					Name:  "abc",
					Price: 2000,
					Details: items.OverviewDetails{
						Status:      items.Ready,
						Stock:       4,
						Size:        3,
						Description: "test",
						Tags:        []string{"aaa", "bbb"},
					},
				},
				Manufacturer: items.ManufacturerDetails{
					Name:            "test",
					StripeAccountId: "test",
					Description:     "test",
					UserId:          "aaa",
				},
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
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err = ManufacturerDB.Register("test", tt.payload, "aaa")
			if err != nil {
				log.Print("error", err.Error())
				t.Errorf("error")
			}
			ItemOverview, err := ItemRepository.GetItemOverview("test")
			if err != nil {
				t.Errorf("error")
			}
			if !reflect.DeepEqual(ItemOverview, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, ItemOverview, tt.want)
			}
			err = ManufacturerDB.Update(tt.updatePayload, "test")
			if err != nil {
				t.Errorf("error")
			}
			ItemOverview, err = ItemRepository.GetItemOverview("test")
			if err != nil {
				t.Errorf("error")
			}
			if !reflect.DeepEqual(ItemOverview, tt.wantUpdated) {
				t.Errorf("%v,got,%v,want%v", tt.name, ItemOverview, tt.wantUpdated)
			}
			err = ManufacturerDB.Delete("test")
			if err != nil {
				t.Errorf("error")
			}

		})
	}
	err = UserDB.Delete("aaa")
	if err != nil {
		t.Errorf("error")
	}
}
func Test_GetItemList(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserDB := users.UserRepository{DB: db}
	ManufacturerDB := TestRepository{DB: db}
	ItemRepository := items.ItemRepository{DB: db}
	Items := []RegisterPayload{
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

	if err = UserDB.Create("aaa"); err != nil {
		t.Errorf("error")
	}
	for _, item := range Items {
		err = ManufacturerDB.Register(item.Name, item, "aaa")
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
		want          []items.Preview
		totalElements int
	}{
		{
			name:      "タグのみで絞り込み",
			pageNum:   1,
			pageSize:  5,
			condition: map[string]interface{}{},
			tags:      []string{"ddd"},
			want: []items.Preview{
				{
					Item_id: "test3",
					Properties: items.PreviewProperties{
						Name:  "test3",
						Price: 4000,
						Details: items.PreviewDetails{
							Status: items.Ready,
						},
					},
				},
				{
					Item_id: "test4",
					Properties: items.PreviewProperties{
						Name:  "test4",
						Price: 4000,
						Details: items.PreviewDetails{
							Status: items.Ready,
						},
					},
				}, {
					Item_id: "test5",
					Properties: items.PreviewProperties{
						Name:  "test5",
						Price: 4000,
						Details: items.PreviewDetails{
							Status: items.Ready,
						},
					},
				}, {
					Item_id: "test6",
					Properties: items.PreviewProperties{
						Name:  "test6",
						Price: 5000,
						Details: items.PreviewDetails{
							Status: items.Ready,
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
			want: []items.Preview{
				{
					Item_id: "test6",
					Properties: items.PreviewProperties{
						Name:  "test6",
						Price: 5000,
						Details: items.PreviewDetails{
							Status: items.Ready,
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
			want: []items.Preview{
				{
					Item_id: "test4",
					Properties: items.PreviewProperties{
						Name:  "test4",
						Price: 4000,
						Details: items.PreviewDetails{
							Status: items.Ready,
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
			want: []items.Preview{
				{
					Item_id: "test3",
					Properties: items.PreviewProperties{
						Name:  "test3",
						Price: 4000,
						Details: items.PreviewDetails{
							Status: items.Ready,
						},
					},
				},
				{
					Item_id: "test4",
					Properties: items.PreviewProperties{
						Name:  "test4",
						Price: 4000,
						Details: items.PreviewDetails{
							Status: items.Ready,
						},
					},
				},
			},
			totalElements: 6,
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			previews, totalElements, err := ItemRepository.GetPreviewList(tt.pageNum, tt.pageSize, tt.condition, tt.tags)
			log.Print("totalElements: ", totalElements)
			log.Print("pre: ", previews)
			if err != nil {
				t.Errorf("error")
			}
			if !reflect.DeepEqual(previews, tt.want) {
				t.Errorf("%v,got,%v,want%v", tt.name, previews, tt.want)
			}
			if totalElements != tt.totalElements {
				t.Errorf("%v,got,%v,want%v", tt.name, totalElements, tt.totalElements)
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
