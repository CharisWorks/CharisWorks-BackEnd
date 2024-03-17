package transaction

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type TransactionRequests struct {
	TransactionRepository IRepository
	CartRepository        cart.IRepository
	CartUtils             cart.IUtils
	StripeRequests        cash.IRequests
	StripeUtils           cash.IUtils
}

func (r TransactionRequests) GetList(userId string) (transactionPreview []TransactionPreview, err error) {
	transactionPreviewList, err := r.TransactionRepository.GetList(userId)
	if err != nil {
		return nil, err
	}
	for _, t := range transactionPreviewList {
		transactionPreview = append(transactionPreview, t)
	}
	return transactionPreview, nil
}

func (r TransactionRequests) GetDetails(userId string, transactionId string) (transactionDetails TransactionDetails, err error) {
	transactionDetails, transactionUserId, _, err := r.TransactionRepository.GetDetails(transactionId)
	if err != nil {
		return transactionDetails, err
	}
	if transactionUserId != userId {
		return transactionDetails, nil
	}
	return transactionDetails, nil
}
func (r TransactionRequests) Purchase(userId string, email string) (clientSecret string, transactionId string, err error) {
	internalCart, err := r.CartRepository.Get(userId)
	if err != nil {
		return clientSecret, transactionId, err
	}
	if len(internalCart) == 0 {
		return clientSecret, transactionId, &utils.InternalError{Message: utils.InternalErrorCartIsEmpty}
	}
	mappedInspectedCart, err := r.CartUtils.Inspect(internalCart)
	if err != nil {
		return clientSecret, transactionId, err
	}
	totalAmount := r.CartUtils.GetTotalAmount(mappedInspectedCart) + 400
	clientSecret, transactionId, err = r.StripeRequests.CreatePaymentintent(userId, totalAmount)
	if err != nil {
		return clientSecret, transactionId, err
	}
	err = r.CartRepository.DeleteAll(userId)
	if err != nil {
		return clientSecret, transactionId, err
	}
	err = r.TransactionRepository.Register(userId, email, transactionId, internalCart)
	if err != nil {
		return clientSecret, transactionId, err
	}
	return clientSecret, transactionId, nil
}
func (r TransactionRequests) PurchaseRefund(stripeTransferId string, stripeTransactionId string) error {
	_, _, transferList, err := r.TransactionRepository.GetDetails(stripeTransactionId)
	if err != nil {
		return err
	}
	for _, t := range transferList {
		if t.transferId == stripeTransferId {
			err = r.StripeUtils.Refund(t.amount, stripeTransactionId, t.stripeAccountId)
			if err != nil {
				return err
			}
			err = r.TransactionRepository.StatusUpdateItems(stripeTransactionId, t.itemId, map[string]interface{}{"status": string(Refund)})
			if err != nil {
				return err
			}
		}
	}
	err = r.TransactionRepository.StatusUpdate(stripeTransactionId, map[string]interface{}{"status": string(Refund)})
	if err != nil {
		return err
	}

	return nil
}
