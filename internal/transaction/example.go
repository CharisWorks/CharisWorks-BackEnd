package transaction

/*
import "github.com/charisworks/charisworks-backend/internal/cash"

type ExampleTransactionRequests struct {
	StripeRequests cash.IRequests
}

func (r ExampleTransactionRequests) GetList(userId string) (*[]TransactionPreview, error) {
	return nil, nil
}
func (r ExampleTransactionRequests) GetDetails(userId string, s string) (*TransactionDetails, error) {
	return nil, nil
}
func (r ExampleTransactionRequests) Purchase(userId string) (*string, error) {
	url, _, err := r.StripeRequests.CreatePaymentintent(userId, 1000)
	if err != nil {
		return nil, err
	}
	return url, nil
}
func (r ExampleTransactionRequests) PurchaseRefund(SessionId string, s string) error {
	return nil
} */

/*
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
*/
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
