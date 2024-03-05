package manufacturer

import (
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (m Repository) Register(itemId string, i ItemRegisterPayload, userId string) error {
	json, err := json.Marshal(i.Details.Tags)
	if err != nil {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	item := utils.Item{
		Id:                 itemId,
		ManufacturerUserId: userId,
		Name:               i.Name,
		Price:              i.Price,
		Status:             string(items.Ready),
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

func (m Repository) Update(i map[string]interface{}, itemId string) error {
	if err := m.DB.Table("items").Where("id = ?", itemId).Updates(i).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
func (m Repository) Delete(itemId string) error {
	if err := m.DB.Table("items").Where("id = ?", itemId).Delete(&utils.Item{}).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}
