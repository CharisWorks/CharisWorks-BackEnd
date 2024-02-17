package cash

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/gin-gonic/gin"
)

type ExampleTransactionUtils struct {
}

func (u ExampleTransactionUtils) GetTotalAmount([]cart.Cart) int64 {
	return 1000
}

func (u ExampleTransactionUtils) InspectCart([]cart.Cart) error {
	return nil
}

type ExampleTransactionRequests struct {
}

func (r ExampleTransactionRequests) GetTransactionList(ctx *gin.Context) ([]TransactionPreview, error) {
	return *new([]TransactionPreview), nil
}
func (r ExampleTransactionRequests) GetTransactionDetails(ctx *gin.Context, TransactionId string) (TransactionDetails, error) {
	return *new(TransactionDetails), nil
}
