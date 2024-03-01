package items

import (
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type ItemDB struct {
	DB *gorm.DB
}

func (r *ItemDB) GetItemOverview(itemId string) (*ItemOverview, error) {
	ItemOverview := new(ItemOverview)
	DBItem := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(DBItem).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}

	tags := new([]string)
	json.Unmarshal([]byte(DBItem.Tags), &tags)
	ItemOverview.Item_id = DBItem.Id
	ItemOverview.Properties = ItemOverviewProperties{
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
func getItemPreview(db *gorm.DB, page int, pageSize int, conditions map[string]interface{}, tags []string) ([]ItemPreview, error) {
	previews := new([]ItemPreview)
	items := new([]utils.Item)
	offset := (page - 1) * pageSize
	query := db.Model(&utils.Item{}).Offset(offset).Limit(pageSize)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	for _, tag := range tags {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}
	err := query.Find(&items).Error
	if err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
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
func (r *ItemDB) GetPreviewList(pageNum int, pageSize int, conditions map[string]interface{}, tags []string) (*[]ItemPreview, error) {
	ItemPreview, err := getItemPreview(r.DB, pageNum, pageSize, conditions, tags)
	if err != nil {
		return nil, err
	}
	return &ItemPreview, nil
}
