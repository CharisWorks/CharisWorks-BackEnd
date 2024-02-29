package manufacturer

import (
	"encoding/json"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type ManufacturerDB struct {
	DB *gorm.DB
}

func (m *ManufacturerDB) RegisterItem(i ItemRegisterPayload, history_item_id int, userId string) error {
	json, err := json.Marshal(i.Details.Tags)
	if err != nil {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	item := utils.Item{
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
	result := m.DB.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *ManufacturerDB) UpdateItem(i map[string]string, history_item_id string, itemId string) error {
	if err := m.DB.Table("items").Where("id = ?", itemId).Update("history_item_id", history_item_id).Updates(i).Error; err != nil {
		return err
	}

	return nil
}
func (m *ManufacturerDB) DeleteItem(itemId string) error {
	if err := m.DB.Table("items").Where("id = ?", itemId).Delete(&utils.Item{}).Error; err != nil {
		return err
	}
	return nil
}
