package admin

import (
	"context"
	"encoding/json"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/utils"
	itempb "github.com/charisworks/charisworks-backend/pkg/grpc"
)

type ItemServiceServer struct {
	itempb.UnimplementedItemServiceServer
}

func (r *ItemServiceServer) Remove(ctx context.Context, req *itempb.RemoveItemRequest) (res *itempb.VoidResponse, err error) {
	db, err := utils.DBInit()
	if err != nil {
		return res, err
	}
	itemId := req.GetItem()
	manufacturerRepository := manufacturer.Repository{DB: db}

	err = manufacturerRepository.Delete(itemId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *ItemServiceServer) All(ctx context.Context, req *itempb.VoidRequest) (res *itempb.AllItemResponse, err error) {
	res = new(itempb.AllItemResponse)
	db, err := utils.DBInit()
	if err != nil {
		return res, err
	}
	itemIds := new([]string)
	itemList := new([]items.Overview)
	db.Table("items").Where("1=1").Select("id").Find(&itemIds)
	itemRepository := items.ItemRepository{DB: db}
	for _, itemId := range *itemIds {
		item, err := itemRepository.GetItemOverview(itemId)
		if err != nil {

			return res, err
		}
		*itemList = append(*itemList, item)
	}
	jsonBytes, err := json.Marshal(itemList)
	if err != nil {
		return res, err
	}
	res.Item = string(jsonBytes)
	return res, nil
}
