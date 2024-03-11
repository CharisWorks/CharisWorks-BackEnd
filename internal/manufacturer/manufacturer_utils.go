package manufacturer

import (
	"github.com/charisworks/charisworks-backend/internal/utils"
)

type ManufacturerUtils struct {
}

func (m ManufacturerUtils) Register(i RegisterPayload) error {
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
func (m ManufacturerUtils) Update(updatePayload UpdatePayload) (map[string]interface{}, error) {
	payload := make(map[string]interface{})
	if updatePayload.Stock > 0 {
		payload["stock"] = updatePayload.Stock
	}
	if updatePayload.Size > 0 {
		payload["size"] = updatePayload.Size
	}
	if len(updatePayload.Description) > 0 {
		payload["description"] = updatePayload.Description
	}
	if len(updatePayload.Tags) > 0 {
		payload["tags"] = updatePayload.Tags
	}
	if updatePayload.Price > 0 {
		payload["price"] = updatePayload.Price
	}
	if len(updatePayload.Status) > 0 {
		payload["status"] = updatePayload.Status
	}
	if len(updatePayload.Name) > 0 {
		payload["name"] = updatePayload.Name
	}
	return payload, nil
}
