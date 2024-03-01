package utils

import "github.com/gin-gonic/gin"

type InternalError struct {
	Message InternalMessage `json:"message"`
}

func (r *InternalError) Error() string {
	return string(r.Message)
}
func Code(i InternalMessage) int {
	statusCode := map[InternalMessage]int{
		InternalErrorInvalidItem:                 400,
		InternalErrorStockOver:                   400,
		InternalErrorInvalidQuantity:             400,
		InternalErrorNoStock:                     400,
		InternalErrorInvalidPayload:              400,
		InternalErrorInvalidQuery:                400,
		InternalErrorInvalidParams:               400,
		InternalErrorInvalidCart:                 400,
		InternalErrorNotFound:                    404,
		InternalErrorDB:                          500,
		InternalErrorInvalidUserRequest:          400,
		InternalErrorManufacturerAlreadyHasBank:  400,
		InternalErrorManufacturerDoesNotHaveBank: 400,
		InternalErrorAccountIsNotSatisfied:       400,
		InternalErrorFromStripe:                  400,
		InternalErrorUnAuthorized:                401,
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
	InternalErrorInvalidPayload     InternalMessage = "The request payload is malformed or contains invalid data."
	InternalErrorInvalidQuery       InternalMessage = "invalid Query"
	InternalErrorInvalidParams      InternalMessage = "invalid Params"
	InternalErrorInvalidCart        InternalMessage = "invalid cart"
	InternalErrorInvalidUserRequest InternalMessage = "invalid user request"
	InternalErrorUnAuthorized       InternalMessage = "unauthorized"
	InternalErrorEmailIsNotVerified InternalMessage = "email is not verified"
)

// DB系
const (
	InternalErrorNotFound InternalMessage = "not found"
	InternalErrorDB       InternalMessage = "DB error"
)

// ストライプ系
const (
	InternalErrorManufacturerAlreadyHasBank  InternalMessage = "manufacturer already has bank"
	InternalErrorManufacturerDoesNotHaveBank InternalMessage = "manufacturer does not have bank"
	InternalErrorAccountIsNotSatisfied       InternalMessage = "account is not satisfied"
	InternalErrorFromStripe                  InternalMessage = "error from stripe"
)
