package manufacturer

import (
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type ManufacturerDB struct {
	DB *gorm.DB
}

func (m ManufacturerDB) RegisterItem(itemId string, i ItemRegisterPayload, history_item_id int, userId string) error {
	json, err := json.Marshal(i.Details.Tags)
	if err != nil {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	item := utils.Item{
		Id:                 itemId,
		ManufacturerUserId: userId,
		HistoryItemId:      history_item_id,
		Name:               i.Name,
		Price:              i.Price,
		Status:             string(items.ItemStatusReady),
		Stock:              i.Details.Stock,
		Size:               i.Details.Size,
		Description:        i.Details.Description,
		Tags:               string(json),
	}
	if err := m.DB.Create(&item).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}

func (m ManufacturerDB) UpdateItem(i map[string]interface{}, history_item_id int, itemId string) error {
	if err := m.DB.Table("items").Where("id = ?", itemId).Update("history_item_id", history_item_id).Updates(i).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (m ManufacturerDB) DeleteItem(itemId string) error {
	if err := m.DB.Table("items").Where("id = ?", itemId).Delete(&utils.Item{}).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
