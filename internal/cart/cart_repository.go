package cart

import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

// ICartDB is an interface for cart database

type CartDB struct {
	DB *gorm.DB
}

func (r CartDB) Get(userId string) (*[]InternalCart, error) {
	InternalCarts := new([]utils.InternalCart)
	resultCart := new([]InternalCart)
	if err := r.DB.Table("carts").
		Select("carts.*, items.*").
		Joins("JOIN items ON carts.item_id = items.id").
		Where("carts.purchaser_user_id = ?", userId).
		Find(&InternalCarts).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	for i, icart := range *InternalCarts {
		cart := new(InternalCart)
		cart.Index = i
		cart.ItemStock = icart.Item.Stock
		cart.Status = items.Status(icart.Item.Status)

		cart.Cart.ItemId = icart.Cart.ItemId
		cart.Cart.Quantity = icart.Cart.Quantity
		cart.Cart.ItemProperties.Name = icart.Item.Name
		cart.Cart.ItemProperties.Price = icart.Item.Price

		*resultCart = append(*resultCart, *cart)
	}
	return resultCart, nil
}

func (r CartDB) Register(userId string, cartRequestPayload CartRequestPayload) error {
	Cart := new(utils.Cart)
	Cart.PurchaserUserId = userId
	Cart.ItemId = cartRequestPayload.ItemId
	Cart.Quantity = cartRequestPayload.Quantity
	if err := r.DB.Create(Cart).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (r CartDB) Update(userId string, cartRequestPayload CartRequestPayload) error {
	if err := r.DB.Table("carts").Where("purchaser_user_id = ?", userId).Where("item_id = ?", cartRequestPayload.ItemId).Update("quantity", cartRequestPayload.Quantity).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}

func (r CartDB) Delete(userId string, itemId string) error {
	log.Print("UserId: ", userId)
	if err := r.DB.Table("carts").Where("purchaser_user_id = ?", userId).Where("item_id = ?", itemId).Delete(utils.Cart{}).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (r CartDB) GetItem(itemId string) (*itemStatus, error) {
	itemDB := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(itemDB).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	itemStatus := new(itemStatus)
	itemStatus.itemStock = itemDB.Stock
	itemStatus.status = items.Status(itemDB.Status)
	return itemStatus, nil
}
