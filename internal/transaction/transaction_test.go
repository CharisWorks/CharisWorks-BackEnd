package transaction

import (
	"log"
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
	manufacturerRequests := manufacturer.Requests{ManufacturerItemRepository: manufacturer.Repository{DB: db}, ManufacturerInspectPayloadUtils: manufacturer.ManufacturerUtils{}, ItemRepository: items.ItemRepository{DB: db}}
	cartRequests := cart.Requests{CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, ItemGetStatus: items.GetStatus{DB: db}}
	transactionRequests := TransactionRequests{TransactionRepository: Repository{DB: db, userRepository: UserRepository}, CartRepository: cartRequests.CartRepository, CartUtils: cartRequests.CartUtils, StripeRequests: cash.Requests{}, StripeUtils: cash.Utils{}}

	Items := []manufacturer.RegisterPayload{
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
	if err = UserRepository.Create("aaa"); err != nil {
		t.Errorf("error")
	}
	if err = UserRepository.UpdateProfile("aaa", map[string]interface{}{
		"display_name":      "test",
		"description":       "test",
		"stripe_account_id": "acct_test",
	}); err != nil {
		t.Errorf("error")
	}
	if err = UserRepository.RegisterAddress("aaa", users.AddressRegisterPayload{
		ZipCode:       "123-4567",
		Address1:      "test",
		Address2:      "test",
		Address3:      "test",
		PhoneNumber:   "00000000000",
		FirstName:     "test",
		LastName:      "test",
		FirstNameKana: "テスト",
		LastNameKana:  "テスト",
	}); err != nil {
		t.Errorf("error")
	}

	for _, item := range Items {
		err = manufacturerRequests.Register(item, "aaa", item.Name)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = manufacturerRequests.Update(manufacturer.UpdatePayload{
			Status: string(items.Available),
		}, "aaa", item.Name)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	carts := []cart.CartRequestPayload{
		{
			ItemId:   "test1",
			Quantity: 2,
		},
		{

			ItemId:   "test2",
			Quantity: 2,
		},
	}

	for _, p := range carts {
		err := cartRequests.Register("aaa", p)
		if err != nil {
			t.Errorf(err.Error())
		}

	}
	clientSecret, err := transactionRequests.Purchase("aaa", "test")
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Print(clientSecret)

	clientSecret, err = transactionRequests.Purchase("aaa", "test")
	if err != nil {
		log.Print(err.Error())
	}
	log.Print(clientSecret)
	db.Table("transactions").Where("purchaser_user_id = ?", "aaa").Delete(utils.Transaction{})
	db.Table("transaction_items").Where("transaction_id = ?", "test").Delete(utils.TransactionItem{})

}
