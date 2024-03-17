package manufacturer

import (
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func TestRegisterItem(t *testing.T) {
	After(t)
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserRepository := users.UserRepository{DB: db}
	userRequests := users.Requests{UserRepository: UserRepository, UserUtils: users.UserUtils{}}
	manufacturerRequests := Requests{ManufacturerItemRepository: TestRepository{DB: db}, ManufacturerInspectPayloadUtils: ManufacturerUtils{}, ItemRepository: items.ItemRepository{DB: db}}
	itemRequests := items.Requests{ItemRepository: items.ItemRepository{DB: db}, ItemUtils: items.ItemUtils{}}

	user_data := []struct {
		userId  string
		profile users.UserProfile
		address users.AddressRegisterPayload
	}{
		{
			userId: "test_user_1",
			profile: users.UserProfile{
				DisplayName:     "test_user_1",
				Description:     "test_user_1",
				StripeAccountId: "acct_1OkjHjPKEl3posmB",
			},
			address: users.AddressRegisterPayload{
				ZipCode:       "123-4567",
				Address1:      "test",
				Address2:      "test",
				Address3:      "test",
				PhoneNumber:   "00000000000",
				FirstName:     "test",
				LastName:      "test",
				FirstNameKana: "テスト",
				LastNameKana:  "テスト",
			},
		},

		{
			userId: "test_user_2",
			profile: users.UserProfile{
				DisplayName:     "test_user_2",
				Description:     "test_user_2",
				StripeAccountId: "",
			},
			address: users.AddressRegisterPayload{
				ZipCode:       "123-4567",
				Address1:      "test",
				Address2:      "test",
				Address3:      "test",
				PhoneNumber:   "00000000000",
				FirstName:     "test",
				LastName:      "test",
				FirstNameKana: "テスト",
				LastNameKana:  "テスト",
			},
		},
		{
			userId: "test_user_3",
			profile: users.UserProfile{
				DisplayName:     "test_user_3",
				Description:     "test_user_3",
				StripeAccountId: "skip",
			},
		},
	}
	for _, u := range user_data {
		if err = userRequests.Create(u.userId); err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
		if u.profile.StripeAccountId != "skip" {
			if err = userRequests.AddressRegister(u.userId, u.address); err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
			if err = userRequests.ProfileUpdate(u.userId, u.profile); err != nil {
				t.Errorf("error")
				After(t)
				return
			}
		}
		db.Table("users").Where("id = ?", u.userId).Updates(map[string]interface{}{"stripe_account_id": u.profile.StripeAccountId})
	}
	Cases := []struct {
		name    string
		userId  string
		payload RegisterPayload
		want    items.Overview
		err     error
	}{
		{
			name:   "正常",
			userId: "test_user_1",
			payload: RegisterPayload{
				Name:  "test_item_1",
				Price: 1000,
				Details: ItemRegisterDetailsPayload{
					Description: "test_item_1",
					Stock:       10,
					Tags:        []string{"aaa", "bbb"},
					Size:        3,
				},
			},
			want: items.Overview{
				Item_id: "test_item_1",
				Properties: items.OverviewProperties{
					Name:  "test_item_1",
					Price: 1000,
					Details: items.OverviewDetails{
						Description: "test_item_1",
						Tags:        []string{"aaa", "bbb"},
						Stock:       10,
						Size:        3,
						Status:      "Ready",
					},
				},
				Manufacturer: items.ManufacturerDetails{
					Name:            "test_user_1",
					StripeAccountId: "acct_1OkjHjPKEl3posmB",
					Description:     "test_user_1",
					UserId:          "test_user_1",
				},
			},
		},
	}
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			err := manufacturerRequests.Register(tt.payload, tt.userId, "test_item_1")
			if err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
			overview, err := itemRequests.GetOverview(tt.payload.Name)
			if err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
			if !reflect.DeepEqual(overview, tt.want) {
				t.Errorf("got %v, want %v", overview, tt.want)
				After(t)
				return
			}

		})
	}
	After(t)
}
func After(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	trdb.Table("transactions").Where("1=1").Delete(utils.Transaction{})
	trdb.Table("transaction_items").Where("1=1").Delete(utils.TransactionItem{})
	db.Table("users").Where("1=1").Delete(utils.User{})
	db.Table("shippings").Where("1=1").Delete(utils.Shipping{})
	db.Table("items").Where("1=1").Delete(utils.Item{})
	db.Table("carts").Where("1=1").Delete(utils.Cart{})

}
