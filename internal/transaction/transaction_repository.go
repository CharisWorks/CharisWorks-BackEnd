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

type Repository struct {
	DB             *gorm.DB
	UserRepository users.IRepository
}

func (r Repository) GetList(userId string) (transactionPreviewList map[string]TransactionPreview, err error) {
	internalTransaction := new([]utils.InternalTransaction)
	if err := r.DB.Table("transactions").
		Select("transactions.*, transaction_items.*").
		Joins("JOIN transaction_items ON transactions.transaction_id = transaction_items.transaction_id").
		Where("transactions.purchaser_user_id = ?", userId).
		Find(&internalTransaction).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	transactionPreviewList = map[string]TransactionPreview{}
	for _, t := range *internalTransaction {
		transactionPreview := new(TransactionPreview)
		transactionItem := new(TransactionItem)
		transactionItem.ItemId = t.TransactionItems.ItemId
		transactionItem.Quantity = t.TransactionItems.Quantity
		transactionItem.Price = t.TransactionItems.Price
		transactionItem.Name = t.TransactionItems.Name
		transactionItem.TransferId = t.TransactionItems.StripeTransferId
		transactionItem.Status = t.TransactionItems.Status
		transactionItem.ManufacturerUserId = t.TransactionItems.ManufacturerUserId
		transactionItem.ManufacturerName = t.TransactionItems.ManufacturerName

		transactionPreview.TransactionId = t.Transaction.TransactionId
		transactionPreview.Status = TransactionStatus(t.Transaction.Status)
		transactionPreview.TransactionAt = t.Transaction.CreatedAt
		transactionPreview.Email = t.Transaction.Email
		_, exist := transactionPreviewList[t.Transaction.TransactionId]
		list := make([]TransactionItem, 0)
		if exist {
			list = transactionPreviewList[t.Transaction.TransactionId].Items
			list = append(list, *transactionItem)
		} else {
			list = append(list, *transactionItem)
		}
		transactionPreview.Items = list
		transactionPreviewList[t.Transaction.TransactionId] = *transactionPreview

	}
	return transactionPreviewList, nil
}
func (r Repository) GetDetails(TransactionId string) (transactionDetails TransactionDetails, userId string, transferList []transfer, err error) {
	internalTransaction := new([]utils.InternalTransaction)
	transactionDetails = *new(TransactionDetails)
	if err := r.DB.Table("transactions").
		Select("transactions.*, transaction_items.*").
		Joins("JOIN transaction_items ON transactions.transaction_id = transaction_items.transaction_id").
		Where("transactions.transaction_id = ?", TransactionId).
		Find(&internalTransaction).Error; err != nil {
		log.Print("DB error: ", err)
		return transactionDetails, "", nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	transferList = make([]transfer, 0)
	itemList := []TransactionItem{}
	for _, t := range *internalTransaction {
		itemList = append(itemList, TransactionItem{
			ItemId:             t.TransactionItems.ItemId,
			Quantity:           t.TransactionItems.Quantity,
			Name:               t.TransactionItems.Name,
			Price:              t.TransactionItems.Price,
			TransferId:         t.TransactionItems.StripeTransferId,
			Status:             t.TransactionItems.Status,
			ManufacturerUserId: t.TransactionItems.ManufacturerUserId,
			ManufacturerName:   t.TransactionItems.ManufacturerName,
		})
		userId = t.Transaction.PurchaserUserId
		transactionDetails.TotalAmount = t.Transaction.TotalAmount
		transactionDetails.TotalPrice = t.Transaction.TotalPrice
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
		transactionDetails.Email = t.Transaction.Email

		tr := transfer{
			amount:          t.TransactionItems.Price * t.TransactionItems.Quantity,
			itemId:          t.TransactionItems.ItemId,
			stripeAccountId: t.TransactionItems.ManufacturerStripeAccountId,
			transferId:      t.TransactionItems.StripeTransferId,
		}
		transferList = append(transferList, tr)

	}
	return transactionDetails, userId, transferList, nil
}

func (r Repository) Register(userId string, email string, transactionId string, internalCartList []cart.InternalCart) error {
	totalPrice := 0
	totalAmount := 0
	transactionItemList := make([]utils.TransactionItem, 0)
	for _, i := range internalCartList {
		t, err := json.Marshal(i.Item.Tags)
		if err != nil {
			return &utils.InternalError{Message: utils.InternalErrorDB}
		}
		transactionItem := utils.TransactionItem{
			TransactionId:               transactionId,
			ItemId:                      i.Cart.ItemId,
			Name:                        i.Item.Name,
			Size:                        i.Item.Size,
			Price:                       i.Item.Price,
			Quantity:                    i.Cart.Quantity,
			Description:                 i.Item.Description,
			Tags:                        string(t),
			ManufacturerUserId:          i.Item.ManufacturerUserId,
			ManufacturerName:            i.Item.ManufacturerName,
			ManufacturerDescription:     i.Item.ManufacturerDescription,
			ManufacturerStripeAccountId: i.Item.ManufacturerStripeId,
			Status:                      string(Pending),
		}
		transactionItemList = append(transactionItemList, transactionItem)
		totalPrice += i.Item.Price * i.Cart.Quantity
		totalAmount += i.Cart.Quantity
	}
	user, err := r.UserRepository.Get(userId)
	if err != nil {
		return err
	}

	if user.UserAddress == *new(users.UserAddress) {
		return &utils.InternalError{Message: utils.InternalErrorAddressIsNotRegistered}
	}
	address := user.UserAddress.Address1 + user.UserAddress.Address2 + user.UserAddress.Address3
	name := user.UserAddress.FirstName + user.UserAddress.LastName
	if err := r.DB.Create(utils.Transaction{
		TransactionId:   transactionId,
		Email:           email,
		PurchaserUserId: userId,
		CreatedAt:       time.Now(),
		ZipCode:         user.UserAddress.ZipCode,
		Address:         address,
		PhoneNumber:     user.UserAddress.PhoneNumber,
		RealName:        name,
		Status:          string(Pending),
		TotalPrice:      totalPrice,
		TotalAmount:     totalAmount,
	}).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}

	for _, i := range transactionItemList {
		if err := r.DB.Create(&i).Error; err != nil {
			log.Print("DB error: ", err)
			return &utils.InternalError{Message: utils.InternalErrorDB}
		}
	}
	return nil
}

func (r Repository) StatusUpdate(transactionId string, conditions map[string]interface{}) error {
	if err := r.DB.Table("transactions").Where("transaction_id = ?", transactionId).Updates(conditions).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}

	return nil
}
func (r Repository) StatusUpdateItems(transactionId string, itemId string, conditions map[string]interface{}) error {
	if err := r.DB.Table("transaction_items").Where("transaction_id = ?", transactionId).Where("item_id", itemId).Updates(conditions).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}

	return nil
}
