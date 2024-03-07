package transaction

import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/transfer"
)

type TransactionRequests struct {
	TransactionRepository ITransactionRepository
	CartRepository        cart.IRepository
	CartUtils             cart.IUtils
	StripeRequests        cash.IStripeRequests
}

func (r TransactionRequests) GetList(userId string) (*[]TransactionPreview, error) {
	transactionPreviewList, err := r.TransactionRepository.GetList(userId)
	if err != nil {
		return nil, err
	}
	transactionPreview := make([]TransactionPreview, len(*transactionPreviewList))
	for _, t := range *transactionPreviewList {
		transactionPreview = append(transactionPreview, t)
	}
	return &transactionPreview, nil
}

func (r TransactionRequests) GetDetails(userId string, transactionId string) (*TransactionDetails, error) {
	transactionDetails, transactionUserId, err := r.TransactionRepository.GetDetails(transactionId)
	if err != nil {
		return nil, err
	}
	if transactionUserId != userId {
		return nil, nil
	}
	return transactionDetails, nil
}
func (r TransactionRequests) Purchase(userId string) (*string, error) {
	internalCart, err := r.CartRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	inspectedCart, err := r.CartUtils.Inspect(*internalCart)
	if err != nil {
		return nil, err
	}
	totalAmount := r.CartUtils.GetTotalAmount(inspectedCart)
	clientSecret, stripeTransactionId, err := r.StripeRequests.CreatePaymentintent(userId, totalAmount)
	if err != nil {
		return nil, err
	}
	err = r.CartRepository.DeleteAll(userId)
	if err != nil {
		return nil, err
	}
	err = r.TransactionRepository.Register(userId, *stripeTransactionId, "transactionId", *internalCart)
	if err != nil {
		return nil, err
	}

	return clientSecret, nil
}

func (r TransactionRequests) PurchaseComplete(stripeTransactionId string) error {

	err := r.TransactionRepository.StatusUpdate(stripeTransactionId, TransactionStatus(Complete))
	if err != nil {
		return err
	}
	return nil
}
func Transfer(amount float64, stripeID string, ItemName string) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthx8UK5v9CdaucJOZj3FwkOHZ8KjDt25IAvplosSab4uybQOyE2Ne6xxxI4Rnh8pWEbYUwPoPG00wvseAHzl"
	log.Print("Transfering... \n amount: ", amount, "\n stripeID: ", stripeID, "\n ItemName: ", ItemName)
	params := &stripe.TransferParams{
		Amount:      stripe.Int64(int64(amount)),
		Currency:    stripe.String(string(stripe.CurrencyJPY)),
		Destination: stripe.String(stripeID),
		Description: stripe.String(ItemName),
	}
	tr, _ := transfer.New(params)
	log.Print(tr.ID)

}
