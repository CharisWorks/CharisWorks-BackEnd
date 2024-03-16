package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InternalError struct {
	Message InternalMessage `json:"message"`
}

func (r *InternalError) Error() string {
	return string(r.Message)
}
func Code(i InternalMessage) int {
	statusCode := map[InternalMessage]int{
		InternalErrorEmailIsNotVerified:          http.StatusUnauthorized,
		InternalErrorInvalidItem:                 http.StatusBadRequest,
		InternalErrorStockOver:                   http.StatusBadRequest,
		InternalErrorInvalidQuantity:             http.StatusBadRequest,
		InternalErrorNoStock:                     http.StatusBadRequest,
		InternalErrorInvalidPayload:              http.StatusBadRequest,
		InternalErrorInvalidQuery:                http.StatusBadRequest,
		InternalErrorInvalidParams:               http.StatusBadRequest,
		InternalErrorInvalidCart:                 http.StatusBadRequest,
		InternalErrorNotFound:                    http.StatusNotFound,
		InternalErrorDB:                          http.StatusInternalServerError,
		InternalErrorInvalidUserRequest:          http.StatusBadRequest,
		InternalErrorManufacturerAlreadyHasBank:  http.StatusBadRequest,
		InternalErrorManufacturerDoesNotHaveBank: http.StatusBadRequest,
		InternalErrorAccountIsNotSatisfied:       http.StatusBadRequest,
		InternalErrorFromStripe:                  http.StatusBadRequest,
		InternalErrorUnAuthorized:                http.StatusUnauthorized,
		InternalErrorIncident:                    http.StatusInternalServerError,
		InternalErrorCartIsEmpty:                 http.StatusBadRequest,
	}
	return statusCode[i]
}
func ReturnErrorResponse(ctx *gin.Context, err error) {
	internalError := err.(*InternalError)
	ctx.JSON(Code(internalError.Message), gin.H{"message": internalError.Message})
}

type InternalMessage string

// 在庫管理系
const (
	InternalErrorInvalidItem     InternalMessage = "invalid item"
	InternalErrorStockOver       InternalMessage = "stock over"
	InternalErrorInvalidQuantity InternalMessage = "invalid quantity"
	InternalErrorNoStock         InternalMessage = "no stock"
)

// リクエスト系
const (
	InternalErrorInvalidPayload         InternalMessage = "The request payload is malformed or contains invalid data."
	InternalErrorInvalidQuery           InternalMessage = "invalid Query"
	InternalErrorInvalidParams          InternalMessage = "invalid Params"
	InternalErrorInvalidCart            InternalMessage = "invalid cart"
	InternalErrorInvalidUserRequest     InternalMessage = "invalid user request"
	InternalErrorUnAuthorized           InternalMessage = "unauthorized"
	InternalErrorEmailIsNotVerified     InternalMessage = "email is not verified"
	InternalErrorAddressIsNotRegistered InternalMessage = "address is not registered"
)

// DB系
const (
	InternalErrorNotFound    InternalMessage = "record not found"
	InternalErrorDB          InternalMessage = "DB error"
	InternalErrorCartIsEmpty InternalMessage = "cart is empty"
)

// ストライプ系
const (
	InternalErrorManufacturerAlreadyHasBank  InternalMessage = "manufacturer already has bank"
	InternalErrorManufacturerDoesNotHaveBank InternalMessage = "manufacturer does not have bank"
	InternalErrorAccountIsNotSatisfied       InternalMessage = "account is not satisfied"
	InternalErrorFromStripe                  InternalMessage = "error from stripe"
)

const (
	InternalErrorIncident InternalMessage = "incident"
)

const (
	InternalErrorR2 InternalMessage = "r2 error"
)
