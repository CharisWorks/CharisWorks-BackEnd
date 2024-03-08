package transaction

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
)

type TransactionRequests struct {
	TransactionRepository ITransactionRepository
	CartRepository        cart.IRepository
	CartUtils             cart.IUtils
	StripeRequests        cash.IRequests
	webhook               Webhook
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
	transactionDetails, transactionUserId, _, err := r.TransactionRepository.GetDetails(transactionId)
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
func (r TransactionRequests) PurchaseRefund(stripeTransferId string, transactionId string, itemId string) error {
	err := r.webhook.PurchaseRefund(stripeTransferId, transactionId)
	if err != nil {
		return err
	}
	err = r.TransactionRepository.StatusUpdateItems(transactionId, itemId, map[string]interface{}{"status": "refunded"})
	if err != nil {
		return err
	}
	err = r.TransactionRepository.StatusUpdate(transactionId, map[string]interface{}{"status": "refunded"})
	if err != nil {
		return err
	}
	return nil
}
