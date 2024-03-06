package cash

import (
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/gin-gonic/gin"
)

type ExampleTransactionRequests struct {
}

func (r ExampleTransactionRequests) GetList(ctx *gin.Context, TransactionDBHistory ITransactionDBHistory) (*[]TransactionPreview, error) {
	return nil, nil
}
func (r ExampleTransactionRequests) GetDetails(ctx *gin.Context) (*TransactionDetails, error) {
	return new(TransactionDetails), nil
}
func (r ExampleTransactionRequests) Create(ctx *gin.Context, CartRequests cart.IRequests, cartRepository cart.IRepository, CartUtils cart.IUtils) error {
	return nil
}

type ExampleTransactionDBHistory struct {
}

func (r ExampleTransactionDBHistory) Create(TransactionDetails TransactionDetails) error {
	return nil
}
func (r ExampleTransactionDBHistory) GetList(UserId string) (*[]TransactionPreview, error) {
	return nil, nil
}
func (r ExampleTransactionDBHistory) GetDetails(TransactionId string) (*TransactionDetails, error) {
	return new(TransactionDetails), nil
}
func (r ExampleTransactionDBHistory) Register(UserId string, TransactionDetails TransactionDetails) (*string, error) {
	return nil, nil
}
func (r ExampleTransactionDBHistory) StatusUpdate(TransactionId string, Status TransactionStatus) error {
	return nil
}

/* type ExampleStripeRequests struct {
}

func (r ExampleStripeRequests) GetClientSecret(ctx *gin.Context, CartRequests cart.ICartRequests, cartRepository cart.IcartRepository, CartUtils cart.ICartUtils, UserId string) (url *string, err error) {
	return nil, nil
}
func (r ExampleStripeRequests) GetRegisterLink(ctx *gin.Context) (url *string, err error) {
	return nil, nil
}
func (r ExampleStripeRequests) GetStripeMypageLink(ctx *gin.Context) (url *string, err error) {
	return nil, nil
}
*/
