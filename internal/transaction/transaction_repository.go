package transaction

import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func (r TransactionRepository) GetList(UserId string) (*map[int]TransactionPreview, error) {
	transactionPreviewList := make(map[int]TransactionPreview)
	internalTransaction := new([]utils.InternalTransaction)
	if err := r.DB.Table("transaction").
		Select("transaction.*, transactionitems.*").
		Joins("JOIN items ON transaction.id = transactionitems.transaction_id").
		Where("transaction.purchaser_user_id = ?", UserId).
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

		transactionPreview.TransactionId = t.Transaction.Id
		transactionPreview.Status = TransactionStatus(t.Transaction.Status)
		transactionPreview.TransactionAt = t.Transaction.CreatedAt
		transaction, exist := transactionPreviewList[t.Transaction.Id]
		if exist {
			transaction.Items = append(transaction.Items, *transactionItem)
		}
		transactionPreviewList[t.Transaction.Id] = transaction
	}
	return &transactionPreviewList, nil
}
func (r TransactionRepository) GetDetails(TransactionId string) (*TransactionDetails, string, error) {
	transactionDetails := new(TransactionDetails)
	internalTransaction := new([]utils.InternalTransaction)
	if err := r.DB.Table("transaction").
		Select("transaction.*, transactionitems.*").
		Joins("JOIN items ON transaction.id = transactionitems.transaction_id").
		Where("transaction.id = ?", TransactionId).
		Find(&internalTransaction).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, "", &utils.InternalError{Message: utils.InternalErrorDB}
	}
	userId := ""
	log.Print(internalTransaction)
	itemList := []TransactionItem{}
	for _, t := range *internalTransaction {
		transactionItem := new(TransactionItem)
		transactionItem.ItemId = t.TransactionItems.ItemId
		transactionItem.Quantity = t.TransactionItems.Quantity
		transactionItem.Price = t.TransactionItems.Price
		transactionItem.Name = t.TransactionItems.Name
		itemList = append(itemList, *transactionItem)
		userId = t.Transaction.PurchaserUserId
		transactionDetails.TransactionId = t.Transaction.Id
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
	}

	return transactionDetails, userId, nil
}

func (r TransactionRepository) Register(TransactionDetails TransactionDetails) error {
	result := r.DB.Create(&TransactionDetails)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (r TransactionRepository) StatusUpdate(TransactionId string, Status TransactionStatus) error {
	if err := r.DB.Table("transactions").Where("id = ?", TransactionId).Update("status", Status).Error; err != nil {
		return err
	}

	return nil
}
