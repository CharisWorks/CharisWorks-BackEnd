package transaction

import (
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/items"
)

type Webhook struct {
	StripeUtils           cash.IUtils
	TransactionRepository IRepository
	ItemUpdater           items.IUpdater
}

func (r Webhook) PurchaseComplete(stripeTransactionId string) error {
	transactionDetails, _, transferList, err := r.TransactionRepository.GetDetails(stripeTransactionId)
	if err != nil {
		return err
	}
	for _, t := range transferList {
		transferId := r.StripeUtils.Transfer(t.amount, t.stripeAccountId, stripeTransactionId)
		if err != nil {
			return err
		}
		err = r.TransactionRepository.StatusUpdateItems(stripeTransactionId, t.itemId, map[string]interface{}{"stripe_transfer_id": transferId, "status": Complete})
		if err != nil {
			return err
		}
	}
	for _, i := range transactionDetails.Items {
		err = r.ItemUpdater.ReduceStock(i.ItemId, i.Quantity)
		if err != nil {
			return err
		}
	}
	err = r.TransactionRepository.StatusUpdate(stripeTransactionId, map[string]interface{}{"status": Complete})
	if err != nil {
		return err
	}

	return nil
}
func (r Webhook) PurchaseFail(stripeTransactionId string) error {
	_, _, transferList, err := r.TransactionRepository.GetDetails(stripeTransactionId)
	if err != nil {
		return err
	}
	for _, t := range transferList {
		err = r.TransactionRepository.StatusUpdateItems(stripeTransactionId, t.itemId, map[string]interface{}{"status": "failed"})
		if err != nil {
			return err
		}
	}
	err = r.TransactionRepository.StatusUpdate(stripeTransactionId, map[string]interface{}{"status": "failed"})
	if err != nil {
		return err
	}

	return nil
}

func (r Webhook) PurchaseCanceled(stripeTransactionId string) error {
	_, _, transferList, err := r.TransactionRepository.GetDetails(stripeTransactionId)
	if err != nil {
		return err
	}
	for _, t := range transferList {
		err = r.TransactionRepository.StatusUpdateItems(stripeTransactionId, t.itemId, map[string]interface{}{"status": string(Cancelled)})
		if err != nil {
			return err
		}
	}
	err = r.TransactionRepository.StatusUpdate(stripeTransactionId, map[string]interface{}{"status": string(Cancelled)})
	if err != nil {
		return err
	}

	return nil
}
