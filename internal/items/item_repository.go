package items

import (
	"encoding/json"
	"log"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func (r ItemRepository) GetItemOverview(itemId string) (*Overview, error) {
	ItemOverview := new(Overview)
	DBItem := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(DBItem).Error; err != nil {
		log.Print("DB error: ", err)
		return nil, &utils.InternalError{Message: utils.InternalErrorDB}
	}

	tags := new([]string)
	json.Unmarshal([]byte(DBItem.Tags), &tags)
	ItemOverview.Item_id = DBItem.Id
	ItemOverview.Properties = OverviewProperties{
		Name:  DBItem.Name,
		Price: DBItem.Price,
		Details: OverviewDetails{
			Status:      Status(DBItem.Status),
			Stock:       DBItem.Stock,
			Size:        DBItem.Size,
			Description: DBItem.Description,
			Tags:        *tags,
		},
	}
	return ItemOverview, nil
}
func getItemPreview(db *gorm.DB, page int, pageSize int, conditions map[string]interface{}, tags []string) ([]Preview, int, error) {
	previews := new([]Preview)
	items := new([]utils.Item)
	offset := (page - 1) * pageSize
	query := db.Model(&utils.Item{})
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	for _, tag := range tags {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}
	totalElements := query.Find(&items).RowsAffected
	err := query.Offset(offset).Limit(pageSize).Find(&items).Error
	if err != nil {
		log.Print("DB error: ", err)
		return nil, 0, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	for _, item := range *items {
		preview := new(Preview)
		preview.Item_id = item.Id
		preview.Properties.Details.Status = Status(item.Status)
		preview.Properties.Name = item.Name
		preview.Properties.Price = item.Price
		*previews = append(*previews, *preview)
	}
	return *previews, int(totalElements), nil
}
func (r ItemRepository) GetPreviewList(pageNum int, pageSize int, conditions map[string]interface{}, tags []string) (*[]Preview, int, error) {
	ItemPreview, totalElements, err := getItemPreview(r.DB, pageNum, pageSize, conditions, tags)
	if err != nil {
		return nil, 0, err
	}
	return &ItemPreview, totalElements, nil
}
