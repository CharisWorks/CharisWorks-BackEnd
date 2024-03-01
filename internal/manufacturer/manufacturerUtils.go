package manufacturer

import (
	"encoding/json"
	"strconv"

	"github.com/charisworks/charisworks-backend/internal/utils"
)

type ManufacturerUtils struct {
}

func (m ManufacturerUtils) InspectRegisterPayload(i ItemRegisterPayload) error {
	if i.Price < 0 {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if i.Name == "" {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if i.Details.Stock < 0 {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if i.Details.Size < 0 {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if i.Details.Description == "" {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}

	return nil
}
func (m ManufacturerUtils) InspectUpdatePayload(i ItemUpdatePayload) (*map[string]string, error) {
	updatepayload := make(map[string]string)
	if i.ItemUpdatePropertiesPayload.Name != nil {
		if *i.ItemUpdatePropertiesPayload.Name != "" {
			updatepayload["name"] = *i.ItemUpdatePropertiesPayload.Name
		} else {
			return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
		}
	}
	if i.ItemUpdatePropertiesPayload.Price != nil {
		if *i.ItemUpdatePropertiesPayload.Price < 0 {
			updatepayload["price"] = strconv.Itoa(*i.ItemUpdatePropertiesPayload.Price)
		} else {
			return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
		}
	}
	if i.ItemUpdatePropertiesPayload.Details != nil {
		if i.ItemUpdatePropertiesPayload.Details.Status != nil {
			if *i.ItemUpdatePropertiesPayload.Details.Status != "" {
				updatepayload["status"] = *i.ItemUpdatePropertiesPayload.Details.Status
			} else {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
		}
		if i.ItemUpdatePropertiesPayload.Details.Stock != nil {
			if *i.ItemUpdatePropertiesPayload.Details.Stock < 0 {
				updatepayload["stock"] = strconv.Itoa(*i.ItemUpdatePropertiesPayload.Details.Stock)
			} else {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
		}
		if i.ItemUpdatePropertiesPayload.Details.Size != nil {
			if *i.ItemUpdatePropertiesPayload.Details.Size < 0 {
				updatepayload["size"] = strconv.Itoa(*i.ItemUpdatePropertiesPayload.Details.Size)
			} else {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
		}
		if i.ItemUpdatePropertiesPayload.Details.Description != nil {
			if *i.ItemUpdatePropertiesPayload.Details.Description != "" {
				updatepayload["description"] = *i.ItemUpdatePropertiesPayload.Details.Description
			} else {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
		}
		if i.ItemUpdatePropertiesPayload.Details.Tags != nil {
			json, err := json.Marshal(i.ItemUpdatePropertiesPayload.Details.Tags)
			if err != nil {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
			updatepayload["tags"] = string(json)
		}
	}
	return &updatepayload, nil
}
