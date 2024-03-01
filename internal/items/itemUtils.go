package items

import (
	"strconv"
	"strings"

	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type ItemUtils struct {
}

func (i ItemUtils) InspectSearchConditions(ctx *gin.Context) (int, int, map[string]interface{}, []string, error) {

	pageNum, _ := utils.GetQuery("page_num", ctx)
	num := new(int)
	if pageNum == nil {
		*num = 1
	} else {
		i, err := strconv.Atoi(*pageNum)
		if err != nil {
			err = &utils.InternalError{Message: utils.InternalErrorInvalidQuery}
			return 0, 0, nil, nil, err
		}
		*num = i
	}

	pageSize, _ := utils.GetQuery("page_size", ctx)
	size := new(int)
	if pageSize == nil {
		*size = 10
	} else {
		i, err := strconv.Atoi(*pageSize)
		if err != nil {
			err = &utils.InternalError{Message: utils.InternalErrorInvalidQuery}
			return 0, 0, nil, nil, err
		}
		*size = i
	}

	tags := make([]string, 0)
	tagString, _ := utils.GetQuery("tags", ctx)
	if tagString == nil {
		tags = strings.Split(*tagString, "+")
	}

	conditions := make(map[string]interface{})
	minPrice, _ := utils.GetQuery("min_price", ctx)
	if minPrice != nil {
		i, err := strconv.Atoi(*minPrice)
		if err != nil {
			err = &utils.InternalError{Message: utils.InternalErrorInvalidQuery}
			return 0, 0, nil, nil, err
		}
		conditions["price >= ?"] = i
	}

	maxPrice, _ := utils.GetQuery("max_price", ctx)
	if maxPrice != nil {
		i, err := strconv.Atoi(*maxPrice)
		if err != nil {
			err = &utils.InternalError{Message: utils.InternalErrorInvalidQuery}
			return 0, 0, nil, nil, err
		}
		conditions["price <= ?"] = i
	}

	return *num, *size, conditions, tags, nil
}
