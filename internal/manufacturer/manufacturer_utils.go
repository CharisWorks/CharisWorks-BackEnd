package manufacturer

import (
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type ManufacturerUtils struct {
}

func (m ManufacturerUtils) InspectRegisterPayload(i ItemRegisterPayload) error {
	if i.Price <= 0 {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if i.Name == "" {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if i.Details.Stock <= 0 {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if i.Details.Size <= 0 {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}
	if i.Details.Description == "" {
		return &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
	}

	return nil
}
func (m ManufacturerUtils) InspectUpdatePayload(i map[string]interface{}) (*map[string]interface{}, error) {
	updatepayload := make(map[string]interface{})
	for k, v := range i {
		if k == "Price" {
			if v.(int) <= 0 {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
			updatepayload[k] = v.(int)
		}
		if k == "Name" {
			if v.(string) == "" {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
			updatepayload[k] = v.(string)
		}
		if k == "Stock" {
			if v.(int) <= 0 {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
			updatepayload[k] = v.(int)
		}
		if k == "Size" {
			if v.(int) <= 0 {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
			updatepayload[k] = v.(int)
		}
		if k == "Description" {
			if v.(string) == "" {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
			updatepayload[k] = v.(string)
		}
		if k == "Tags" {
			if len(v.([]string)) == 0 {
				return nil, &utils.InternalError{Message: utils.InternalErrorInvalidPayload}
			}
			updatepayload[k] = v.([]string)
		}

	}

	return &updatepayload, nil
}
