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
	After(t)

	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserRepository := users.UserRepository{DB: db}
	userRequests := users.Requests{UserRepository: UserRepository, UserUtils: users.UserUtils{}}
	manufacturerRequests := manufacturer.Requests{ManufacturerItemRepository: manufacturer.Repository{DB: db}, ManufacturerInspectPayloadUtils: manufacturer.ManufacturerUtils{}, ItemRepository: items.ItemRepository{DB: db}}
	cartRequests := cart.Requests{CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, ItemGetStatus: items.GetStatus{DB: db}}
	transactionRequests := TransactionRequests{TransactionRepository: Repository{DB: trdb, UserRepository: UserRepository}, CartRepository: cartRequests.CartRepository, CartUtils: cartRequests.CartUtils, StripeRequests: cash.Requests{}, StripeUtils: cash.Utils{}}
	webhook := Webhook{StripeUtils: cash.Utils{}, TransactionRepository: Repository{DB: trdb, UserRepository: UserRepository}, ItemUpdater: items.Updater{DB: db}}
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
			After(t)
			return
		}
		err = manufacturerRequests.Update(manufacturer.UpdatePayload{
			Status: string(items.Available),
		}, "test_user_1", item.Name)
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
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
			name:   "住所登録がまだ",
			userId: "test_user_3",
			carts: []cart.CartRequestPayload{
				{

					ItemId:   "test2",
					Quantity: 1,
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorAddressIsNotRegistered},
		},
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
				Email:       "hoge@example.com",
				TotalAmount: 4,
				TotalPrice:  10000,
				UserAddress: TransactionAddress{
					ZipCode:     "123-4567",
					Address:     "testtesttest",
					PhoneNumber: "000-0000-0000",
					RealName:    "testtest",
				},
				Items: []TransactionItem{
					{
						ItemId:             "test1",
						Quantity:           2,
						Name:               "test1",
						Price:              2000,
						Status:             string(Pending),
						ManufacturerUserId: "test_user_1",
						ManufacturerName:   "test_user_1",
					},
					{
						ItemId:             "test2",
						Quantity:           2,
						Name:               "test2",
						Price:              3000,
						Status:             string(Pending),
						ManufacturerUserId: "test_user_1",
						ManufacturerName:   "test_user_1",
					},
				},
			},
			transactionPreview: []TransactionPreview{
				{
					TransactionId: "test",
					Email:         "hoge@example.com",
					Items: []TransactionItem{
						{
							ItemId:             "test1",
							Quantity:           2,
							Name:               "test1",
							Price:              2000,
							ManufacturerUserId: "test_user_1",
							ManufacturerName:   "test_user_1",
						},
						{
							ItemId:             "test2",
							Quantity:           2,
							Name:               "test2",
							Price:              3000,
							ManufacturerUserId: "test_user_1",
							ManufacturerName:   "test_user_1",
						},
					},
				},
			},
		},
		{
			name:   "在庫不足",
			userId: "test_user_1",
			carts: []cart.CartRequestPayload{
				{
					ItemId:   "test2",
					Quantity: 2,
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorStockOver},
		},
		{
			name:   "在庫なし",
			userId: "test_user_1",
			carts: []cart.CartRequestPayload{
				{
					ItemId:   "test1",
					Quantity: 2,
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorNoStock},
		},
		{
			name:   "カートが空",
			userId: "test_user_1",
			err:    &utils.InternalError{Message: utils.InternalErrorCartIsEmpty},
		},
		{
			name:   "ユーザーが存在しない",
			userId: "test_user_4",
			carts: []cart.CartRequestPayload{
				{
					ItemId:   "test2",
					Quantity: 1,
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorDB},
		},
		{
			name:   "商品が存在しない",
			userId: "test_user_1",
			carts: []cart.CartRequestPayload{
				{
					ItemId:   "test3",
					Quantity: 1,
				},
			},
			err: &utils.InternalError{Message: utils.InternalErrorDB},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			items := new([]utils.Item)
			db.Table("items").Where("1=1").Find(&items)

			for _, cart := range c.carts {
				err := cartRequests.Register(c.userId, cart)
				if err != nil {
					if err.Error() != c.err.Error() {
						t.Errorf("got %v, want %v", err, c.err)
						After(t)
					}
					return
				}
			}
			_, transactionId, err := transactionRequests.Purchase(c.userId, "hoge@example.com")
			if err != nil {
				if err.Error() != c.err.Error() {
					t.Errorf("got %v, want %v", err, c.err)
					After(t)
				}
				return

			}
			transactionDetails, err := transactionRequests.GetDetails(c.userId, transactionId)
			if err != nil {
				if err.Error() != c.err.Error() {
					t.Errorf("got %v, want %v", err, c.err)
					After(t)
				}
				return

			}
			c.transactionDetails.Status = Pending
			c.transactionDetails.TransactionAt = transactionDetails.TransactionAt
			tr := []TransactionItem{}
			for _, t := range transactionDetails.Items {
				t.Status = string(Pending)
				tr = append(tr, t)
			}
			c.transactionDetails.Items = tr
			c.transactionDetails.TransactionId = transactionDetails.TransactionId
			if !reflect.DeepEqual(transactionDetails, c.transactionDetails) {
				t.Errorf("got %v, want %v", transactionDetails, c.transactionDetails)
			}
			transactionPreview, err := transactionRequests.GetList("test_user_1")
			if err != nil {
				if err.Error() != c.err.Error() {
					t.Errorf("got %v, want %v", err, c.err)
					After(t)
				}
				return
			}
			c.transactionPreview[0].TransactionId = transactionId
			c.transactionPreview[0].Status = Pending
			c.transactionPreview[0].TransactionAt = transactionDetails.TransactionAt
			c.transactionPreview[0].Items[0].Status = string(Pending)
			c.transactionPreview[0].Items[1].Status = string(Pending)

			if !reflect.DeepEqual(transactionPreview, c.transactionPreview) {
				t.Errorf("got %v, want %v", transactionPreview, c.transactionPreview)
			}
			_, err = webhook.PurchaseComplete(transactionId)
			if err != nil {
				if err.Error() != c.err.Error() {
					t.Errorf("got %v, want %v", err, c.err)
					After(t)
				}

				return
			}
			transactionDetails, err = transactionRequests.GetDetails(c.userId, transactionId)
			if err != nil {
				if err.Error() != c.err.Error() {
					t.Errorf("got %v, want %v", err, c.err)
					After(t)
				}
				return

			}
			c.transactionDetails.Status = Complete
			titems := []TransactionItem{}
			for i, t := range c.transactionDetails.Items {
				t.Status = string(Complete)
				t.TransferId = transactionDetails.Items[i].TransferId
				titems = append(titems, t)
			}
			c.transactionDetails.Items = titems
			if !reflect.DeepEqual(transactionDetails, c.transactionDetails) {
				t.Errorf("got %v, want %v", transactionDetails, c.transactionDetails)
			}
			transactionPreview, err = transactionRequests.GetList("test_user_1")
			if err != nil {
				if err.Error() != c.err.Error() {
					t.Errorf("got %v, want %v", err, c.err)
					After(t)
				}
				return
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

		})
	}
	log.Print("test finished")
	for _, user := range user_data {
		db.Table("users").Where("id = ?", user.userId).Delete(utils.User{})
	}

	After(t)

}
func Test_Transaction_Cancelled(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserRepository := users.UserRepository{DB: db}
	userRequests := users.Requests{UserRepository: UserRepository, UserUtils: users.UserUtils{}}
	manufacturerRequests := manufacturer.Requests{ManufacturerItemRepository: manufacturer.Repository{DB: db}, ManufacturerInspectPayloadUtils: manufacturer.ManufacturerUtils{}, ItemRepository: items.ItemRepository{DB: db}}
	cartRequests := cart.Requests{CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, ItemGetStatus: items.GetStatus{DB: db}}
	transactionRequests := TransactionRequests{TransactionRepository: Repository{DB: trdb, UserRepository: UserRepository}, CartRepository: cartRequests.CartRepository, CartUtils: cartRequests.CartUtils, StripeRequests: cash.Requests{}, StripeUtils: cash.Utils{}}
	webhook := Webhook{StripeUtils: cash.Utils{}, TransactionRepository: Repository{DB: trdb, UserRepository: UserRepository}, ItemUpdater: items.Updater{DB: db}}
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
			After(t)
			return
		}
		err = manufacturerRequests.Update(manufacturer.UpdatePayload{
			Status: string(items.Available),
		}, "test_user_1", item.Name)
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
	}

	t.Run("testing", func(t *testing.T) {
		cartRequests.Register("test_user_1", cart.CartRequestPayload{
			ItemId:   "test1",
			Quantity: 2,
		})
		cartRequests.Register("test_user_1", cart.CartRequestPayload{
			ItemId:   "test2",
			Quantity: 2,
		})
		_, transactionId, err := transactionRequests.Purchase("test_user_1", "hoge@example.com")
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
		_, err = webhook.PurchaseComplete(transactionId)
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
		err = webhook.PurchaseCanceled(transactionId)
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
		transactionDetails, err := transactionRequests.GetDetails("test_user_1", transactionId)
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
		log.Print(transactionDetails)
		failedTransaction := TransactionDetails{
			TransactionAt: transactionDetails.TransactionAt,
			Email:         "hoge@example.com",
			TransactionId: transactionId,
			Status:        Cancelled,
			TotalAmount:   4,
			TotalPrice:    10000,
			Items: []TransactionItem{
				{
					ItemId:             "test1",
					Quantity:           2,
					Name:               "test1",
					Price:              2000,
					Status:             string(Cancelled),
					TransferId:         transactionDetails.Items[0].TransferId,
					ManufacturerUserId: "test_user_1",
					ManufacturerName:   "test1",
				},
				{
					ItemId:             "test2",
					Quantity:           2,
					Name:               "test2",
					Price:              3000,
					Status:             string(Cancelled),
					TransferId:         transactionDetails.Items[1].TransferId,
					ManufacturerUserId: "test_user_1",
					ManufacturerName:   "test1",
				},
			},
			UserAddress: TransactionAddress{
				ZipCode:     "123-4567",
				Address:     "testtesttest",
				PhoneNumber: "000-0000-0000",
				RealName:    "testtest",
			},
		}
		if !reflect.DeepEqual(transactionDetails, failedTransaction) {
			t.Errorf("got %v, want %v", transactionDetails, failedTransaction)
		}

	})

	log.Print("test finished")
	for _, user := range user_data {
		db.Table("users").Where("id = ?", user.userId).Delete(utils.User{})
	}

	After(t)

}

/*
	 func Test_Transaction_Refund(t *testing.T) {
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
				After(t)
				return
			}
			err = manufacturerRequests.Update(manufacturer.UpdatePayload{
				Status: string(items.Available),
			}, "test_user_1", item.Name)
			if err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
		}

		t.Run("testing", func(t *testing.T) {
			cartRequests.Register("test_user_1", cart.CartRequestPayload{
				ItemId:   "test1",
				Quantity: 2,
			})
			cartRequests.Register("test_user_1", cart.CartRequestPayload{
				ItemId:   "test2",
				Quantity: 2,
			})
			_, transactionId, err := transactionRequests.Purchase("test_user_1")
			if err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
			err = webhook.PurchaseComplete(transactionId)
			if err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
			transactionDetails, err := transactionRequests.GetDetails("test_user_1", transactionId)
			if err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
			err = transactionRequests.PurchaseRefund(transactionDetails.Items[0].TransferId, transactionId)
			if err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
			transactionDetails, err = transactionRequests.GetDetails("test_user_1", transactionId)
			if err != nil {
				t.Errorf(err.Error())
				After(t)
				return
			}
			log.Print(transactionDetails)
		})

		log.Print("test finished")
		for _, user := range user_data {
			db.Table("users").Where("id = ?", user.userId).Delete(utils.User{})
		}

		//db.Table("transactions").Where("purchaser_user_id = ?", "aaa").Delete(utils.Transaction{})
		//db.Table("transaction_items").Where("transaction_id = ?", tr.TransactionId).Delete(utils.TransactionItem{})
		After(t)

}
*/
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
