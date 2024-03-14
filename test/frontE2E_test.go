package e2e

import (
	"log"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/admin"
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func TestE2E(t *testing.T) {

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
	transactionRequests := transaction.TransactionRequests{TransactionRepository: transaction.Repository{DB: trdb, UserRepository: UserRepository}, CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, StripeRequests: cash.Requests{CartRequests: cartRequests, UserRequests: userRequests}, StripeUtils: cash.Utils{}}
	webhook := transaction.Webhook{StripeUtils: cash.Utils{}, TransactionRepository: transaction.Repository{DB: trdb, UserRepository: UserRepository}, ItemUpdater: items.Updater{DB: db}}

	user_data := []struct {
		userId  string
		profile users.UserProfile
		address users.AddressRegisterPayload
	}{

		{
			userId: "WQElviFCW3TEV77prNZB7Q2TwGt2",
			profile: users.UserProfile{
				DisplayName:     "つっちー",
				Description:     "つっちーだよ☆",
				StripeAccountId: "acct_1OkjHjPKEl3posmB",
			},
			address: users.AddressRegisterPayload{
				ZipCode:       "123-4567",
				Address1:      "宮城県仙台市青葉区",
				Address2:      "通町",
				Address3:      "通町マンション209",
				PhoneNumber:   "02280561422",
				FirstName:     "土屋",
				LastName:      "徳三郎",
				FirstNameKana: "ツチヤ",
				LastNameKana:  "トクサブロウ",
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
		err = manufacturerRequests.Register(item, "WQElviFCW3TEV77prNZB7Q2TwGt2", item.Name)
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
		err = manufacturerRequests.Update(manufacturer.UpdatePayload{
			Status: string(items.Available),
		}, "WQElviFCW3TEV77prNZB7Q2TwGt2", item.Name)
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
	}
	carts := []struct {
		userId string
		cart   cart.CartRequestPayload
	}{
		{
			userId: "WQElviFCW3TEV77prNZB7Q2TwGt2",
			cart: cart.CartRequestPayload{
				ItemId:   "test1",
				Quantity: 1,
			},
		},
	}
	for _, c := range carts {
		err = cartRequests.Register(c.userId, c.cart)
		if err != nil {
			t.Errorf(err.Error())
			After(t)
			return
		}
	}

	clientSecret, transactionId, err := transactionRequests.Purchase("WQElviFCW3TEV77prNZB7Q2TwGt2", "cowatanabe26@gmail.com")
	if err != nil {
		t.Errorf(err.Error())
		After(t)
		return
	}
	log.Print("clientSecret: ", clientSecret)
	log.Print("transactionId: ", transactionId)
	transactionDetails, err := webhook.PurchaseComplete(transactionId)
	if err != nil {
		t.Errorf("got %v", err)
		After(t)
	}
	admin.SendPurchasedEmail(transactionDetails)
	log.Print("test finished")
}
func After(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	trdb, _ := utils.HistoryDBInitTest()
	trdb.Table("transactions").Where("1=1").Delete(utils.Transaction{})
	trdb.Table("transaction_items").Where("1=1").Delete(utils.TransactionItem{})
	db.Table("users").Where("1=1").Delete(utils.User{})
	db.Table("shippings").Where("1=1").Delete(utils.Shipping{})
	db.Table("items").Where("1=1").Delete(utils.Item{})
	db.Table("carts").Where("1=1").Delete(utils.Cart{})

}
