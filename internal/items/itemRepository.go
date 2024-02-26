package items

import (
	"github.com/charisworks/charisworks-backend/internal/utils"
	"gorm.io/gorm"
)

type ItemDB struct {
	DB *gorm.DB
}

func (r *ItemDB) GetItemOverview(itemId string) (*utils.Item, error) {
	ItemOverview := new(ItemOverview)
	DBItem := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(DBItem).Error; err != nil {
		return nil, err
	}
	ItemOverview.Item_id = DBItem.Id
	ItemOverview.Properties = &ItemOverviewProperties{
		Name:  &DBItem.Name,
		Price: &DBItem.Price,
		Details: &ItemOverviewDetails{
			Status:      ItemStatus(DBItem.Status),
			Stock:       &DBItem.Stock,
			Size:        &DBItem.Size,
			Description: &DBItem.Description,
			Tags:        &DBItem.Tags,
		},
	}
	return DBItem, nil
}

func (r *ItemDB) GetPreviewList(keywords *[]string, page *string, manufacturer *string) (*[]ItemPreview, error) {

	return nil, nil
}
