package transaction

import (
	"encoding/json"
	"log"
	"time"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB             *gorm.DB
	userRepository users.IRepository
}

func (r TransactionRepository) GetList(userId string) (*map[string]TransactionPreview, error) {
	transactionPreviewList := make(map[string]TransactionPreview)
	internalTransaction := new([]utils.InternalTransaction)
	if err := r.DB.Table("transactions").
		Select("transactions.*, transaction_items.*").
		Joins("JOIN transaction_items ON transactions.transaction_id = transaction_items.transaction_id").
		Where("transactions.purchaser_user_id = ?", userId).
		Find(&internalTransaction).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	log.Print(internalTransaction)
	for _, t := range *internalTransaction {
		transactionPreview := new(TransactionPreview)
		transactionItem := new(TransactionItem)
		transactionItem.ItemId = t.TransactionItems.ItemId
		transactionItem.Quantity = t.TransactionItems.Quantity
		transactionItem.Price = t.TransactionItems.Price
		transactionItem.Name = t.TransactionItems.Name

		transactionPreview.TransactionId = t.Transaction.TransactionId
		transactionPreview.Status = TransactionStatus(t.Transaction.Status)
		transactionPreview.TransactionAt = t.Transaction.CreatedAt
		transaction, exist := transactionPreviewList[t.Transaction.TransactionId]
		if exist {
			transaction.Items = append(transaction.Items, *transactionItem)
		}
		transactionPreviewList[t.Transaction.TransactionId] = transaction
	}
	return &transactionPreviewList, nil
}
func (r TransactionRepository) GetDetails(TransactionId string) (transactionDetails *TransactionDetails, userId string, transferList []transfer, err error) {
	internalTransaction := new([]utils.InternalTransaction)
	if err := r.DB.Table("transactions").
		Select("transactions.*, transaction_items.*").
		Joins("JOIN transaction_items ON transactions.id = transaction_items.transaction_id").
		Where("transactions.id = ?", TransactionId).
		Find(&internalTransaction).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, "", nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	transactionDetails = new(TransactionDetails)
	transferList = make([]transfer, 0)
	log.Print(internalTransaction)
	itemList := []TransactionItem{}
	for _, t := range *internalTransaction {
		itemList = append(itemList, TransactionItem{
			ItemId:     t.TransactionItems.ItemId,
			Quantity:   t.TransactionItems.Quantity,
			Name:       t.TransactionItems.Name,
			Price:      t.TransactionItems.Price,
			TransferId: t.TransactionItems.StripeTransferId,
			Status:     t.TransactionItems.Status,
		})
		userId = t.Transaction.PurchaserUserId
		transactionDetails.TransactionId = TransactionId
		transactionDetails.Status = TransactionStatus(t.Transaction.Status)
		transactionDetails.TransactionAt = t.Transaction.CreatedAt
		transactionDetails.Items = itemList
		transactionDetails.TrackingId = t.Transaction.TrackingId
		transactionDetails.UserAddress = TransactionAddress{
			ZipCode:     t.Transaction.ZipCode,
			Address:     t.Transaction.Address,
			PhoneNumber: t.Transaction.PhoneNumber,
			RealName:    t.Transaction.RealName,
		}

		tr := transfer{
			amount:          t.TransactionItems.Price * t.TransactionItems.Quantity,
			itemId:          t.TransactionItems.ItemId,
			stripeAccountId: t.TransactionItems.ManufacturerStripeAccountId,
		}
		transferList = append(transferList, tr)

	}

	return transactionDetails, userId, transferList, nil
}

func (r TransactionRepository) Register(userId string, stripeTransactionId string, transactionId string, internalCartList []cart.InternalCart) error {
	totalPrice := 0
	totalAmount := 0
	transactionItemList := make([]utils.TransactionItem, 0)
	for _, i := range internalCartList {
		t, err := json.Marshal(i.Item.Tags)
		if err != nil {
			return &utils.InternalError{Message: utils.InternalErrorDB}
		}
		transactionItem := utils.TransactionItem{
			TransactionId:           transactionId,
			ItemId:                  i.Cart.ItemId,
			Name:                    i.Item.Name,
			Price:                   i.Item.Price,
			Quantity:                i.Cart.Quantity,
			Description:             i.Item.Description,
			Tags:                    string(t),
			ManufacturerUserId:      i.Item.ManufacturerUserId,
			ManufacturerName:        i.Item.ManufacturerName,
			ManufacturerDescription: i.Item.ManufacturerDescription,
		}
		transactionItemList = append(transactionItemList, transactionItem)
		totalPrice += i.Item.Price * i.Cart.Quantity
		totalAmount += i.Cart.Quantity
	}
	user, err := r.userRepository.Get(userId)
	if err != nil {
		return err
	}
	address := user.UserAddress.Address1 + user.UserAddress.Address2 + user.UserAddress.Address3
	name := user.UserAddress.FirstName + user.UserAddress.LastName
	if err := r.DB.Create(utils.Transaction{
		TransactionId:       transactionId,
		PurchaserUserId:     userId,
		CreatedAt:           time.Now(),
		ZipCode:             user.UserAddress.ZipCode,
		Address:             address,
		PhoneNumber:         user.UserAddress.PhoneNumber,
		RealName:            name,
		Status:              string(Pending),
		StripeTransactionId: stripeTransactionId,
		TotalPrice:          totalPrice,
		TotalAmount:         totalAmount,
	}).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	for _, i := range transactionItemList {
		if err := r.DB.Create(&i).Error; err != nil {
			log.Print("DB error: ", err)
		}
	}
	return nil
}

func (r TransactionRepository) StatusUpdate(stripeTransactionId string, conditions map[string]interface{}) error {
	if err := r.DB.Table("transactions").Where("stripe_transaction_id = ?", stripeTransactionId).Updates(conditions).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}

	return nil
}
func (r TransactionRepository) StatusUpdateItems(stripeTransactionId string, itemId string, conditions map[string]interface{}) error {
	if err := r.DB.Where("stripe_transaction_id = ?", stripeTransactionId).Where("item_id", itemId).Updates(conditions).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}

	return nil
}
