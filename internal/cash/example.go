package cash

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

type ExampleTransactionRequests struct {
}

func (r ExampleTransactionRequests) GetTransactionList(ctx *gin.Context, CartRequests cart.ICartRequests, CartDB cart.ICartDB, CartUtils cart.ICartUtils, ItemDB items.IItemDB) (*[]TransactionPreview, error) {
	return nil, nil
}
func (r ExampleTransactionRequests) GetTransactionDetails(ctx *gin.Context, TransactionId string) (TransactionDetails, error) {
	return *new(TransactionDetails), nil
}
func (r ExampleTransactionRequests) CreateTransaction(ctx *gin.Context, CartRequests cart.ICartRequests, CartDB cart.ICartDB, CartUtils cart.ICartUtils, userId string) error {
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
