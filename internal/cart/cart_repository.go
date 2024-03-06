package cart

import (
	"log"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

// IcartRepository is an interface for cart database

type Repository struct {
	DB *gorm.DB
}

func (r Repository) Get(UserId string) (*[]InternalCart, error) {
	InternalCarts := new([]utils.InternalCart)
	resultCart := new([]InternalCart)
	if err := r.DB.Table("carts").
		Select("carts.*, items.*").
		Joins("JOIN items ON carts.item_id = items.id").
		Where("carts.purchaser_user_id = ?", UserId).
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

func (r Repository) Register(UserId string, CartRequestPayload CartRequestPayload) error {
	Cart := new(utils.Cart)
	Cart.PurchaserUserId = UserId
	Cart.ItemId = CartRequestPayload.ItemId
	Cart.Quantity = CartRequestPayload.Quantity
	if err := r.DB.Create(Cart).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (r Repository) Update(UserId string, CartRequestPayload CartRequestPayload) error {
	if err := r.DB.Table("carts").Where("purchaser_user_id = ?", UserId).Where("item_id = ?", CartRequestPayload.ItemId).Update("quantity", CartRequestPayload.Quantity).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}

func (r Repository) Delete(UserId string, itemId string) error {
	log.Print("UserId: ", UserId)
	if err := r.DB.Table("carts").Where("purchaser_user_id = ?", UserId).Where("item_id = ?", itemId).Delete(utils.Cart{}).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (r Repository) GetItem(itemId string) (*itemStatus, error) {
	ItemRepository := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(ItemRepository).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	itemStatus := new(itemStatus)
	itemStatus.itemStock = ItemRepository.Stock
	itemStatus.status = items.Status(ItemRepository.Status)
	return itemStatus, nil
}
