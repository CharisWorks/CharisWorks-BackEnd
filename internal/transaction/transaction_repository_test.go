package transaction

import (
	"log"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/cart"
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
	ManufacturerDB := manufacturer.Repository{DB: db}
	cartRepository := cart.Repository{DB: db}
	transactionRepository := Repository{DB: db, userRepository: UserRepository}
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
		"stripe_account_id": "test",
	}); err != nil {
		t.Errorf("error")
	}
	if err = UserRepository.RegisterAddress("aaa", users.AddressRegisterPayload{
		ZipCode:       "123-4567",
		Address1:      "test",
		Address2:      "test",
		Address3:      "test",
		PhoneNumber:   "test",
		FirstName:     "test",
		LastName:      "test",
		FirstNameKana: "test",
		LastNameKana:  "test",
	}); err != nil {
		t.Errorf("error")
	}

	for _, item := range Items {
		if err != nil {
			t.Errorf("error")
		}
		err = ManufacturerDB.Update(map[string]interface{}{"status": items.Available}, item.Name)
		if err != nil {
			t.Errorf("error")
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
		err := cartRepository.Register("aaa", p)
		if err != nil {
			t.Errorf(err.Error())
		}

	}
	cart, err := cartRepository.Get("aaa")
	if err != nil {
		t.Errorf(err.Error())
	}
	transactionRepository.Register("aaa", "test", "test", *cart)
	transaction, err := transactionRepository.GetList("aaa")
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Print("transaction: ", transaction)

	transactioDetails, user, transfer, err := transactionRepository.GetDetails("test")
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Print("transactionDetails: ", transactioDetails, "\n user: ", user, "\ntransfer", transfer)

	db.Table("transactions").Where("purchaser_user_id = ?", "aaa").Delete(utils.Transaction{})
	db.Table("transaction_items").Where("transaction_id = ?", "test").Delete(utils.TransactionItem{})

	for _, item := range Items {
		err = ManufacturerDB.Delete(item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}
	err = UserRepository.Delete("aaa")
	if err != nil {
		t.Errorf("error")
	}
}
