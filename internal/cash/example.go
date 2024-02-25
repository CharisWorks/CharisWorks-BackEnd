package cash

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/gin-gonic/gin"
)

type ExampleTransactionRequests struct {
}

func (r ExampleTransactionRequests) GetTransactionList(ctx *gin.Context, TransactionDBHistory ITransactionDBHistory, userId string) (*[]TransactionPreview, error) {
	return nil, nil
}
func (r ExampleTransactionRequests) GetTransactionDetails(ctx *gin.Context, TransactionId string) (*TransactionDetails, error) {
	return new(TransactionDetails), nil
}
func (r ExampleTransactionRequests) CreateTransaction(ctx *gin.Context, CartRequests cart.ICartRequests, CartDB cart.ICartDB, CartUtils cart.ICartUtils, userId string) error {
	return nil
}

type ExampleTransactionDBHistory struct {
}

func (r ExampleTransactionDBHistory) CreateTransaction(TransactionDetails TransactionDetails) error {
	return nil
}
func (r ExampleTransactionDBHistory) GetTransactionList(userId string) (*[]TransactionPreview, error) {
	return nil, nil
}
func (r ExampleTransactionDBHistory) GetTransactionDetails(TransactionId string) (*TransactionDetails, error) {
	return new(TransactionDetails), nil
}
func (r ExampleTransactionDBHistory) RegisterTransaction(userId string, TransactionDetails TransactionDetails) (*string, error) {
	return nil, nil
}
func (r ExampleTransactionDBHistory) TransactionStatusUpdate(TransactionId string, Status TransactionStatus) error {
	return nil
}

/* type ExampleStripeRequests struct {
}

func (r ExampleStripeRequests) GetClientSecret(ctx *gin.Context, CartRequests cart.ICartRequests, CartDB cart.ICartDB, CartUtils cart.ICartUtils, userId string) (url *string, err error) {
	return nil, nil
}
func (r ExampleStripeRequests) GetRegisterLink(ctx *gin.Context) (url *string, err error) {
	return nil, nil
}
func (r ExampleStripeRequests) GetStripeMypageLink(ctx *gin.Context) (url *string, err error) {
	return nil, nil
}
*/
