package items

import (
	"encoding/json"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type ItemDB struct {
	DB *gorm.DB
}

func (r *ItemDB) GetItemOverview(itemId int) (*ItemOverview, error) {
	ItemOverview := new(ItemOverview)
	DBItem := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(DBItem).Error; err != nil {
		return nil, err
	}

	tags := new([]string)
	json.Unmarshal([]byte(DBItem.Tags), &tags)

	ItemOverview.Item_id = DBItem.Id
	ItemOverview.Properties = &ItemOverviewProperties{
		Name:  DBItem.Name,
		Price: DBItem.Price,
		Details: ItemOverviewDetails{
			Status:      ItemStatus(DBItem.Status),
			Stock:       DBItem.Stock,
			Size:        DBItem.Size,
			Description: DBItem.Description,
			Tags:        *tags,
		},
	}
	return ItemOverview, nil
}
func getItemPreview(db *gorm.DB, page int, pageSize int, conditions map[string]interface{}) ([]ItemPreview, error) {
	previews := new([]ItemPreview)
	items := new([]utils.Item)
	offset := (page - 1) * pageSize

	query := db.Model(&utils.Item{}).Offset(offset).Limit(pageSize)
	for key, value := range conditions {
		query = query.Where(key, value)
	}

	err := query.Find(&items).Error
	if err != nil {
		return nil, err
	}
	for _, item := range *items {
		preview := new(ItemPreview)
		preview.Item_id = item.Id
		preview.Properties.Details.Status = ItemStatus(item.Status)
		preview.Properties.Name = item.Name
		preview.Properties.Price = item.Price
		*previews = append(*previews, *preview)
	}
	return *previews, nil
}
func (r *ItemDB) GetPreviewList(pageNum *int, pageSize *int, conditions map[string]interface{}) (*[]ItemPreview, error) {
	if pageNum == nil {
		*pageNum = 1
	}
	if pageSize == nil {
		*pageSize = 20
	}
	ItemPreview, err := getItemPreview(r.DB, *pageNum, *pageSize, conditions)
	if err != nil {
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return &ItemPreview, nil
}
