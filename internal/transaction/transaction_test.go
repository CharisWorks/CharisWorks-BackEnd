package transaction

import (
	"log"
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func Test_Transaction(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}

	UserRepository := users.UserRepository{DB: db}
	userRequests := users.Requests{UserRepository: UserRepository, UserUtils: users.UserUtils{}}
	manufacturerRequests := manufacturer.Requests{ManufacturerItemRepository: manufacturer.Repository{DB: db}, ManufacturerInspectPayloadUtils: manufacturer.ManufacturerUtils{}, ItemRepository: items.ItemRepository{DB: db}}
	cartRequests := cart.Requests{CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, ItemGetStatus: items.GetStatus{DB: db}}
	transactionRequests := TransactionRequests{TransactionRepository: Repository{DB: db, UserRepository: UserRepository}, CartRepository: cartRequests.CartRepository, CartUtils: cartRequests.CartUtils, StripeRequests: cash.Requests{}, StripeUtils: cash.Utils{}}
	webhook := Webhook{StripeUtils: cash.Utils{}, TransactionRepository: Repository{DB: db, UserRepository: UserRepository}, ItemUpdater: items.Updater{DB: db}}
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
	}
	for _, u := range user_data {
		if err = userRequests.Create(u.userId); err != nil {
			t.Errorf("error")
		}
		if err = userRequests.AddressRegister(u.userId, u.address); err != nil {
			t.Errorf("error")
		}
		if err = userRequests.ProfileUpdate(u.userId, u.profile); err != nil {
			t.Errorf("error")
		}
		db.Table("users").Where("id = ?", u.userId).Updates(map[string]interface{}{"stripe_account_id": u.profile.StripeAccountId})
	}
	item_data := []manufacturer.RegisterPayload{
		{
			Name:  "test1",
			Price: 2000,
			Details: manufacturer.ItemRegisterDetailsPayload{
				Stock:       2,
				Size:        3,
				Description: "test",
				Tags:        []string{"aaa", "bbb"},
			}},
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
	for _, item := range item_data {
		err = manufacturerRequests.Register(item, "test_user_1", item.Name)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = manufacturerRequests.Update(manufacturer.UpdatePayload{
			Status: string(items.Available),
		}, "test_user_1", item.Name)
		if err != nil {
			t.Errorf(err.Error())
		}

	}
	cases := []struct {
		name               string
		userId             string
		carts              []cart.CartRequestPayload
		transactionDetails TransactionDetails
		transactionPreview []TransactionPreview
		err                error
	}{
		{
			name:   "正常",
			userId: "test_user_1",
			carts: []cart.CartRequestPayload{
				{
					ItemId:   "test1",
					Quantity: 2,
				},
				{

					ItemId:   "test2",
					Quantity: 2,
				},
			},
			transactionDetails: TransactionDetails{
				UserAddress: TransactionAddress{
					ZipCode:     "123-4567",
					Address:     "testtesttest",
					PhoneNumber: "000-0000-0000",
					RealName:    "testtest",
				},
				Items: []TransactionItem{
					{
						ItemId:   "test1",
						Quantity: 2,
						Name:     "test1",
						Price:    2000,
						Status:   string(Pending),
					},
					{
						ItemId:   "test2",
						Quantity: 2,
						Name:     "test2",
						Price:    3000,
						Status:   string(Pending),
					},
				},
			},
			transactionPreview: []TransactionPreview{
				{
					TransactionId: "test",
					Items: []TransactionItem{
						{
							ItemId:   "test1",
							Quantity: 2,
							Name:     "test1",
							Price:    2000,
						},
						{
							ItemId:   "test2",
							Quantity: 2,
							Name:     "test2",
							Price:    3000,
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		for _, cart := range c.carts {
			err := cartRequests.Register(c.userId, cart)
			if err != nil {
				t.Errorf(err.Error())
			}
		}
		_, transactionId, err := transactionRequests.Purchase(c.userId)
		if err != nil {
			t.Errorf(err.Error())
		}
		transactionDetails, err := transactionRequests.GetDetails(c.userId, transactionId)
		if err != nil {
			t.Errorf(err.Error())
		}
		c.transactionDetails.Status = Pending
		c.transactionDetails.TransactionAt = transactionDetails.TransactionAt
		c.transactionDetails.Items[0].Status = string(Pending)
		c.transactionDetails.Items[1].Status = string(Pending)
		c.transactionDetails.TransactionId = transactionDetails.TransactionId
		if !reflect.DeepEqual(transactionDetails, c.transactionDetails) {
			t.Errorf("got %v, want %v", transactionDetails, c.transactionDetails)
		}
		transactionPreview, err := transactionRequests.GetList("test_user_1")
		if err != nil {
			t.Errorf(err.Error())
		}
		c.transactionPreview[0].TransactionId = transactionId
		c.transactionPreview[0].Status = Pending
		c.transactionPreview[0].TransactionAt = transactionDetails.TransactionAt
		c.transactionPreview[0].Items[0].Status = string(Pending)
		c.transactionPreview[0].Items[1].Status = string(Pending)

		if !reflect.DeepEqual(transactionPreview, c.transactionPreview) {
			t.Errorf("got %v, want %v", transactionPreview, c.transactionPreview)
		}
		err = webhook.PurchaseComplete(transactionId)
		if err != nil {
			t.Errorf(err.Error())
		}
		transactionDetails, err = transactionRequests.GetDetails(c.userId, transactionId)
		if err != nil {
			t.Errorf(err.Error())
		}
		c.transactionDetails.Status = Complete
		c.transactionDetails.Items[0].Status = string(Complete)
		c.transactionDetails.Items[1].Status = string(Complete)
		c.transactionDetails.Items[0].TransferId = transactionDetails.Items[0].TransferId
		c.transactionDetails.Items[1].TransferId = transactionDetails.Items[1].TransferId
		if !reflect.DeepEqual(transactionDetails, c.transactionDetails) {
			t.Errorf("got %v, want %v", transactionDetails, c.transactionDetails)
		}
		transactionPreview, err = transactionRequests.GetList("test_user_1")
		if err != nil {
			t.Errorf(err.Error())
		}
		c.transactionPreview[0].Status = Complete
		c.transactionPreview[0].TransactionAt = transactionDetails.TransactionAt
		c.transactionPreview[0].Items[0].Status = string(Complete)
		c.transactionPreview[0].Items[1].Status = string(Complete)
		c.transactionPreview[0].Items[0].TransferId = transactionPreview[0].Items[0].TransferId
		c.transactionPreview[0].Items[1].TransferId = transactionPreview[0].Items[1].TransferId
		if !reflect.DeepEqual(transactionPreview, c.transactionPreview) {
			t.Errorf("got %v, want %v", transactionPreview, c.transactionPreview)
		}

		db.Table("transactions").Where("purchaser_user_id = ?", c.userId).Delete(utils.Transaction{})
		db.Table("transaction_items").Where("transaction_id = ?", transactionId).Delete(utils.TransactionItem{})
	}
	for _, user := range user_data {
		db.Table("users").Where("id = ?", user.userId).Delete(utils.User{})
	}

	//db.Table("transactions").Where("purchaser_user_id = ?", "aaa").Delete(utils.Transaction{})
	//db.Table("transaction_items").Where("transaction_id = ?", tr.TransactionId).Delete(utils.TransactionItem{})
	items := new([]utils.Item)
	db.Table("items").Find(&items)
	log.Print("items:", items)
	err = UserRepository.Delete("aaa")
	if err != nil {
		t.Errorf(err.Error())
	}

}
