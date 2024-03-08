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
	UserDB := users.UserRepository{DB: db}
	ManufacturerDB := manufacturer.Repository{DB: db}
	cartRepository := cart.Repository{DB: db}
	transactionRepository := TransactionRepository{DB: db, userRepository: UserDB}
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
	if err = UserDB.RegisterAddress("aaa", users.AddressRegisterPayload{
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
		err = ManufacturerDB.Register(item.Name, item, "aaa")
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
	for _, item := range Items {
		err = ManufacturerDB.Delete(item.Name)
		if err != nil {
			t.Errorf("error")
		}
	}

	transactioDetails, user, transfer, err := transactionRepository.GetDetails("test")
	log.Print("transactioDetails: ", transactioDetails, user, transfer)
	db.Table("transactions").Where("purchaser_user_id = ?", "aaa").Delete(utils.Transaction{})
	db.Table("transaction_items").Where("transaction_id = ?", "test").Delete(utils.TransactionItem{})
	err = UserDB.Delete("aaa")
	if err != nil {
		t.Errorf("error")
	}
}
