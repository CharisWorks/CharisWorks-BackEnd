package utils

type InternalError struct {
	Message InternalErrorMessage
}

func (e *InternalError) Error() string {
	return string(e.Message)
}

type InternalErrorMessage string

const (
	InternalErrorInvalidItem     InternalErrorMessage = "invalid item"
	InternalErrorStockOver       InternalErrorMessage = "stock over"
	InternalErrorInvalidQuantity InternalErrorMessage = "invalid quantity"
	InternalErrorNoStock         InternalErrorMessage = "no stock"
	InternalErrorInvalidPayload  InternalErrorMessage = "invaild payload"
	InternalErrorInvalidQuery    InternalErrorMessage = "invalid Query"
	InternalErrorInvalidParams   InternalErrorMessage = "invalid Params"
	InternalErrorInvalidCart     InternalErrorMessage = "invalid cart"
	InternalErrorNotFound        InternalErrorMessage = "not found"
)
