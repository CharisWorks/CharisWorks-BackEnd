package cart

import (
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

// IcartRepository is an interface for cart database

type Repository struct {
	DB *gorm.DB
}

func (r Repository) Get(UserId string) (internalCart []InternalCart, err error) {
	InternalCarts := new([]utils.InternalCart)
	internalCart = *new([]InternalCart)
	if err := r.DB.Table("carts").
		Select("carts.*, items.*,users.*").
		Joins("JOIN items ON carts.item_id = items.id").
		Joins("JOIN users ON items.manufacturer_user_id = users.id").
		Where("carts.purchaser_user_id = ?", UserId).
		Find(&InternalCarts).Error; err != nil {
		log.Print("DB error: ", err)
		if err.Error() == "record not found" {
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
		} else {
			err = &utils.InternalError{Message: utils.InternalErrorDB}
		}
		return internalCart, err
	}

	for i, icart := range *InternalCarts {
		cart := new(InternalCart)
		cart.Index = i
		cart.ItemStock = icart.Item.Stock
		cart.Status = items.Status(icart.Item.Status)
		tags := new([]string)
		err := json.Unmarshal([]byte(icart.Item.Tags), &tags)
		if err != nil {
			log.Print("DB error: ", err)
			return internalCart, &utils.InternalError{Message: utils.InternalErrorDB}
		}
		cart.Cart.ItemId = icart.Cart.ItemId
		cart.Cart.Quantity = icart.Cart.Quantity
		cart.Cart.ItemProperties.Name = icart.Item.Name
		cart.Cart.ItemProperties.Price = icart.Item.Price
		cart.Cart.ItemProperties.Details.Status = ItemStatus(cart.Status)
		cart.Item.Name = icart.Item.Name
		cart.Item.Price = icart.Item.Price
		cart.Item.Description = icart.Item.Description
		cart.Item.Size = icart.Item.Size
		cart.Item.Tags = *tags
		cart.Item.ManufacturerDescription = icart.User.Description
		cart.Item.ManufacturerName = icart.User.DisplayName
		cart.Item.ManufacturerUserId = icart.User.Id
		cart.Item.ManufacturerStripeId = icart.User.StripeAccountId
		internalCart = append(internalCart, *cart)
	}
	return internalCart, nil
}

func (r Repository) Register(UserId string, CartRequestPayload CartRequestPayload) error {
	Cart := new(utils.Cart)
	Cart.PurchaserUserId = UserId
	Cart.ItemId = CartRequestPayload.ItemId
	Cart.Quantity = CartRequestPayload.Quantity
	if err := r.DB.Create(Cart).Error; err != nil {
		log.Print("DB error: ", err)
		if err.Error() == "record not found" {
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
		} else {
			err = &utils.InternalError{Message: utils.InternalErrorDB}
		}
		return err
	}
	return nil
}
func (r Repository) Update(UserId string, CartRequestPayload CartRequestPayload) error {
	if err := r.DB.Table("carts").Where("purchaser_user_id = ?", UserId).Where("item_id = ?", CartRequestPayload.ItemId).Update("quantity", CartRequestPayload.Quantity).Error; err != nil {
		log.Print("DB error: ", err)
		if err.Error() == "record not found" {
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
		} else {
			err = &utils.InternalError{Message: utils.InternalErrorDB}
		}
		return err
	}
	return nil
}

func (r Repository) Delete(UserId string, itemId string) error {
	if err := r.DB.Table("carts").Where("purchaser_user_id = ?", UserId).Where("item_id = ?", itemId).Delete(utils.Cart{}).Error; err != nil {
		log.Print("DB error: ", err)
		if err.Error() == "record not found" {
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
		} else {
			err = &utils.InternalError{Message: utils.InternalErrorDB}
		}
		return err
	}
	return nil
}
func (r Repository) DeleteAll(UserId string) error {
	if err := r.DB.Table("carts").Where("purchaser_user_id = ?", UserId).Delete(utils.Cart{}).Error; err != nil {
		log.Print("DB error: ", err)
		if err.Error() == "record not found" {
			err = &utils.InternalError{Message: utils.InternalErrorNotFound}
		} else {
			err = &utils.InternalError{Message: utils.InternalErrorDB}
		}
		return err
	}
	return nil
}
