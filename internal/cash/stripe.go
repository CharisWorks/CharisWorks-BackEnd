package cash

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/account"
	"github.com/stripe/stripe-go/v76/accountlink"
	"github.com/stripe/stripe-go/v76/loginlink"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

type StripeRequests struct {
}

func (StripeRequests StripeRequests) GetRegisterLink(ctx *gin.Context) (*string, error) {

	email := ctx.MustGet("UserEmail").(string)
	User := ctx.MustGet("User").(user.User)

	Account, err := GetAccount(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return nil, err
	}

	if Account.PayoutsEnabled {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "アカウントが存在しています。"})
		return nil, nil
	}
	params := &stripe.AccountParams{

		Capabilities: &stripe.AccountCapabilitiesParams{
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
			BankTransferPayments: &stripe.AccountCapabilitiesBankTransferPaymentsParams{
				Requested: stripe.Bool(true),
			},
		},
		Country:      stripe.String("JP"),
		Email:        stripe.String(email),
		Type:         stripe.String(*stripe.String(string(stripe.AccountTypeExpress))),
		BusinessType: stripe.String(*stripe.String(string(stripe.AccountBusinessTypeIndividual))),
		Individual: &stripe.PersonParams{
			FirstNameKanji: stripe.String(User.UserAddress.FirstName),
			FirstNameKana:  stripe.String(User.UserAddress.FirstNameKana),
			LastNameKanji:  stripe.String(User.UserAddress.LastName),
			LastNameKana:   stripe.String(User.UserAddress.LastNameKana),
			Email:          stripe.String(email),
			Phone:          stripe.String(User.UserAddress.PhoneNumber),
		},
		BusinessProfile: &stripe.AccountBusinessProfileParams{
			MCC:                stripe.String("5699"),
			URL:                stripe.String("charis.works/user/profile/" + User.UserId),
			ProductDescription: stripe.String("this is an account of manufacturer for charis works"),
		},
	}

	a, err := account.New(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return nil, err
	}

	URL, err := CreateAccountLink(ctx, a.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return nil, err
	}
	return URL, nil

}

func CreateAccountLink(ctx *gin.Context, StripeAccountId string) (*string, error) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String(StripeAccountId),
		RefreshURL: stripe.String("http://localhost:3000"),
		ReturnURL:  stripe.String("http://localhost:3000"),
		Type:       stripe.String("account_onboarding"),
		Collect:    stripe.String("eventually_due"),
	}
	result, err := accountlink.New(params)
	if err != nil {
		return nil, err
	}
	return &result.URL, nil
}
func (StripeRequests StripeRequests) GetStripeMypageLink(ctx *gin.Context) (*string, error) {
	Account, err := GetAccount(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return nil, err
	}
	if !Account.PayoutsEnabled {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "口座が登録されていません。"})
		return nil, nil
	}
	params := &stripe.LoginLinkParams{Account: &Account.ID}
	result, err := loginlink.New(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}
	return &result.URL, nil
}

func GetAccount(ctx *gin.Context) (*stripe.Account, error) {
	params := &stripe.AccountParams{}
	StripeAccountId := ctx.MustGet("User").(user.User).Manufacturer.StripeAccountId
	if *StripeAccountId == "" {
		result := new(stripe.Account)
		result.PayoutsEnabled = false
		return result, nil
	}
	result, err := account.GetByID(*StripeAccountId, params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (StripeRequests StripeRequests) GetClientSecret(ctx *gin.Context, CartRequests cart.ICartRequests, CartDB cart.ICartDB, CartUtils cart.ICartUtils) (*string, error) {
	Carts, err := CartDB.GetCart(ctx.MustGet("UserId").(string))
	if err != nil {
		return nil, err
	}
	InspectedCart, err := CartUtils.InspectCart(*Carts)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}
	totalAmount := int64(CartUtils.GetTotalAmount(InspectedCart))

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(totalAmount), //合計金額を算出する関数をインジェクト
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		// In the latest version of the API, specifying the `automatic_payment_methods` parameter is optional because Stripe enables its functionality by default.
		PaymentMethodTypes: []*string{stripe.String("card"), stripe.String("konbini")},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Printf("pi.New: %v", err)
		return nil, err
	}
	return &pi.ClientSecret, nil
}
