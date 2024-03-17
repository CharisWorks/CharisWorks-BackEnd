/*
	 import (
		"github.com/gin-gonic/gin"

)

type ExampleManufacturerRequests struct {
}

	func (m ExampleManufacturerRequests) RegisterItem(ctx *gin.Context) error {
		return nil
	}

	func (m ExampleManufacturerRequests) UpdateItem(ctx *gin.Context) error {
		return nil
	}

	func (m ExampleManufacturerRequests) DeleteItem(ctx *gin.Context) error {
		return nil
	}
*/
package manufacturer

import (
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type TestRepository struct {
	DB *gorm.DB
}

func (r TestRepository) Register(itemId string, i RegisterPayload, userId string) error {
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
	if err := r.DB.Create(&item).Error; err != nil {
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

func (r TestRepository) Update(i map[string]interface{}, itemId string) error {
	if err := r.DB.Table("items").Where("id = ?", itemId).Updates(i).Error; err != nil {
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
func (r TestRepository) Delete(itemId string) error {
	if err := r.DB.Table("items").Where("id = ?", itemId).Delete(&utils.Item{}).Error; err != nil {
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
