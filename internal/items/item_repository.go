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

func (r ItemRepository) GetItemOverview(itemId string) (overview Overview, err error) {
	DBItem := new(utils.InternalItem)
	if err := r.DB.Table("items").Select("items.*, users.*").Joins("JOIN users ON items.manufacturer_user_id = users.id").Where("items.id = ?", itemId).First(DBItem).Error; err != nil {
		log.Print("DB error: ", err)
		return overview, &utils.InternalError{Message: utils.InternalErrorDB}
	}

	tags := new([]string)
	json.Unmarshal([]byte(DBItem.Item.Tags), &tags)
	overview = Overview{
		Item_id: DBItem.Item.Id,
		Properties: OverviewProperties{
			Name:  DBItem.Item.Name,
			Price: DBItem.Item.Price,
			Details: OverviewDetails{
				Status:      Status(DBItem.Item.Status),
				Stock:       DBItem.Item.Stock,
				Size:        DBItem.Item.Size,
				Description: DBItem.Item.Description,
				Tags:        *tags,
			},
		},
		Manufacturer: ManufacturerDetails{
			Name:            DBItem.User.DisplayName,
			StripeAccountId: DBItem.User.StripeAccountId,
			Description:     DBItem.User.Description,
			UserId:          DBItem.User.Id,
		},
	}

	return overview, nil
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
func (r ItemRepository) GetPreviewList(pageNum int, pageSize int, conditions map[string]interface{}, tags []string) ([]Preview, int, error) {
	ItemPreview, totalElements, err := getItemPreview(r.DB, pageNum, pageSize, conditions, tags)
	if err != nil {
		return nil, 0, err
	}
	return ItemPreview, totalElements, nil
}

type GetStatus struct {
	DB *gorm.DB
}

func (r GetStatus) GetItem(itemId string) (status ItemStatus, err error) {
	ItemRepository := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(ItemRepository).Error; err != nil {
		log.Print("DB error: ", err)
		return status, &utils.InternalError{Message: utils.InternalErrorDB}
	}
	status.Stock = ItemRepository.Stock
	status.Status = Status(ItemRepository.Status)
	return status, nil
}

type Updater struct {
	DB *gorm.DB
}

func (r Updater) ReduceStock(itemId string, Quantity int) error {
	ItemRepository := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(ItemRepository).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	ItemRepository.Stock -= Quantity
	if err := r.DB.Table("items").Where("id = ?", itemId).Updates(ItemRepository).Error; err != nil {
		log.Print("DB error: ", err)
		return &utils.InternalError{Message: utils.InternalErrorDB}
	}
	return nil
}

func (r Updater) StatusUpdate(itemId string, State Status) {
	ItemRepository := new(utils.Item)
	if err := r.DB.Table("items").Where("id = ?", itemId).First(ItemRepository).Error; err != nil {
		log.Print("DB error: ", err)
	}
	ItemRepository.Status = string(State)
	if err := r.DB.Table("items").Where("id = ?", itemId).Updates(ItemRepository).Error; err != nil {
		log.Print("DB error: ", err)
	}
}
