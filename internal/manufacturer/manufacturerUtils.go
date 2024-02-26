package manufacturer

import (
	"encoding/json"

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
	if i.ItemUpdatePropertiesPayload.Name != nil && *i.ItemUpdatePropertiesPayload.Name != "" {
		updatepayload["name"] = *i.ItemUpdatePropertiesPayload.Name
	}
	if i.ItemUpdatePropertiesPayload.Price != nil && *i.ItemUpdatePropertiesPayload.Price < 0 {
		updatepayload["price"] = string(rune(*i.ItemUpdatePropertiesPayload.Price))
	}
	if i.ItemUpdatePropertiesPayload.Details != nil {
		if i.ItemUpdatePropertiesPayload.Details.Status != nil && *i.ItemUpdatePropertiesPayload.Details.Status != "" {
			updatepayload["status"] = *i.ItemUpdatePropertiesPayload.Details.Status
		}
		if i.ItemUpdatePropertiesPayload.Details.Stock != nil && *i.ItemUpdatePropertiesPayload.Details.Stock < 0 {
			updatepayload["stock"] = string(rune(*i.ItemUpdatePropertiesPayload.Details.Stock))
		}
		if i.ItemUpdatePropertiesPayload.Details.Size != nil && *i.ItemUpdatePropertiesPayload.Details.Size < 0 {
			updatepayload["size"] = string(rune(*i.ItemUpdatePropertiesPayload.Details.Size))
		}
		if i.ItemUpdatePropertiesPayload.Details.Description != nil && *i.ItemUpdatePropertiesPayload.Details.Description != "" {
			updatepayload["description"] = *i.ItemUpdatePropertiesPayload.Details.Description
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
