package transaction

import (
	"reflect"
	"testing"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

func Test_Transaction_Repository(t *testing.T) {
	db, err := utils.DBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	trdb, err := utils.HistoryDBInitTest()
	if err != nil {
		t.Errorf("error")
	}
	UserRepository := users.UserRepository{DB: db}
	manufacturerRequests := manufacturer.Requests{ManufacturerItemRepository: manufacturer.Repository{DB: db}, ManufacturerInspectPayloadUtils: manufacturer.ManufacturerUtils{}, ItemRepository: items.ItemRepository{DB: db}}
	manufacturerRepository := manufacturer.Repository{DB: db}
	cartRequests := cart.Requests{CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, ItemGetStatus: items.GetStatus{DB: db}}
	transactionRepository := Repository{DB: trdb, UserRepository: UserRepository}
	cartRepository := cart.Repository{DB: db}
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
		err = manufacturerRepository.Register(item.Name, item, "aaa")
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
	c, err := cartRepository.Get("aaa")
	if err != nil {
		t.Errorf(err.Error())
	}
	transactionRepository.Register("aaa", "hoge@example.com", "test", c)
	transaction, err := transactionRepository.GetList("aaa")
	if err != nil {
		t.Errorf(err.Error())
	}
	list := TransactionPreview{
		TransactionId: "test",
		Items: []TransactionItem{
			{
				ItemId:     "test1",
				Quantity:   2,
				Name:       "test1",
				Price:      2000,
				TransferId: "",
				Status:     "Pending",
			},
			{
				ItemId:     "test2",
				Quantity:   2,
				Name:       "test2",
				Price:      3000,
				TransferId: "",
				Status:     "Pending",
			},
		},
	}
	if !reflect.DeepEqual(transaction["test"].Items, list.Items) {
		t.Errorf("got %v, want %v", transaction["test"].Items, list.Items)
	}
	transactionDetails, _, _, err := transactionRepository.GetDetails("test")
	if err != nil {
		t.Errorf(err.Error())
	}
	details := TransactionDetails{
		TransactionId: "test",
		Email:         "hoge@example.com",
		TotalAmount:   4,
		TotalPrice:    10000,
		TrackingId:    "",
		UserAddress: TransactionAddress{
			ZipCode:     "123-4567",
			Address:     "testtesttest",
			PhoneNumber: "00000000000",
			RealName:    "testtest",
		},
		Items: []TransactionItem{
			{
				ItemId:     "test1",
				Quantity:   2,
				Name:       "test1",
				Price:      2000,
				TransferId: "",
				Status:     "Pending",
			},
			{
				ItemId:     "test2",
				Quantity:   2,
				Name:       "test2",
				Price:      3000,
				TransferId: "",
				Status:     "Pending",
			},
		},
		Status: "Pending",
	}
	details.TransactionAt = transactionDetails.TransactionAt
	if !reflect.DeepEqual(transactionDetails, details) {
		t.Errorf("got %v, want %v", transactionDetails, details)
	}
	err = transactionRepository.StatusUpdate("test", map[string]interface{}{"status": "completed", "tracking_id": "test"})
	if err != nil {
		t.Errorf(err.Error())
	}
	err = transactionRepository.StatusUpdateItems("test", "test1", map[string]interface{}{"stripe_transfer_id": "test", "status": "completed"})
	if err != nil {
		t.Errorf(err.Error())
	}
	transactionDetails, _, _, err = transactionRepository.GetDetails("test")
	if err != nil {
		t.Errorf(err.Error())
	}
	details.Status = "completed"
	details.TrackingId = "test"
	details.Items[0].Status = "completed"
	details.Items[0].TransferId = "test"
	if !reflect.DeepEqual(transactionDetails, details) {
		t.Errorf("got %v, want %v", transactionDetails, details)
	}
	trdb.Table("transactions").Where("purchaser_user_id = ?", "aaa").Delete(utils.Transaction{})
	trdb.Table("transaction_items").Where("transaction_id = ?", "test").Delete(utils.TransactionItem{})

	for _, item := range Items {
		err = manufacturerRequests.Delete(item.Name, "aaa")
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	err = UserRepository.Delete("aaa")
	if err != nil {
		t.Errorf("error")
	}
}
