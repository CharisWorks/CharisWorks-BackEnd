package cash

import (
	"github.com/gin-gonic/gin"
)

type TransactionRequests struct {
}

func (r TransactionRequests) GetTransactionList(ctx *gin.Context, TransactionDBHistory ITransactionDBHistory, UserId string) (*[]TransactionPreview, error) {

	return nil, nil
}
