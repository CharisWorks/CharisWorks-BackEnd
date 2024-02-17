package cash

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/gin-gonic/gin"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/account"
	"github.com/stripe/stripe-go/v76/accountlink"
	"github.com/stripe/stripe-go/v76/accountsession"
	"github.com/stripe/stripe-go/v76/loginlink"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func HandleCreatePaymentIntent(ctx *gin.Context) {

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1000), //合計金額を算出する関数をインジェクト
		Currency: stripe.String(string(stripe.CurrencyJPY)),
		// In the latest version of the API, specifying the `automatic_payment_methods` parameter is optional because Stripe enables its functionality by default.
		PaymentMethodTypes: []*string{stripe.String("card"), stripe.String("konbini")},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		log.Printf("pi.New: %v", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"clientSecret": pi.ClientSecret})

}

func CreateStripeAccount(ctx *gin.Context) {

	email := ctx.MustGet("UserEmail").(string)
	User := ctx.MustGet("User").(*user.User)

	params := &stripe.AccountParams{}
	result, err := account.GetByID(ctx.MustGet("Stripe_Account_Id").(string), params)
	if err != nil {
		return
	}
	log.Print(result.PayoutsEnabled)
	if result.PayoutsEnabled {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "アカウントが存在しています。"})
		return
	}
	params = &stripe.AccountParams{

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

	a, _ := account.New(params)
	ctx.Set("Stripe_Account_Id", a.ID)
	URL := CreateAccountLink(ctx)
	ctx.JSON(http.StatusOK, gin.H{"url": URL})
}
func CreateAccountLink(ctx *gin.Context) string {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String(ctx.MustGet("Stripe_Account_Id").(string)),
		RefreshURL: stripe.String("http://localhost:3000"),
		ReturnURL:  stripe.String("http://localhost:3000"),
		Type:       stripe.String("account_onboarding"),
		Collect:    stripe.String("eventually_due"),
	}
	result, err := accountlink.New(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result.URL)
	ctx.JSON(http.StatusOK, gin.H{"url": result.URL})
	return result.URL
}
func GetMypage(ctx *gin.Context) {

	params := &stripe.LoginLinkParams{Account: stripe.String(ctx.MustGet("Stripe_Account_Id").(string))}
	result, err := loginlink.New(params)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"url": result.URL})
}
func CreateAccountSession(ctx *gin.Context) {
	params := &stripe.AccountSessionParams{
		Account: stripe.String(ctx.MustGet("Stripe_Account_Id").(string)),
		Components: &stripe.AccountSessionComponentsParams{
			AccountOnboarding: &stripe.AccountSessionComponentsAccountOnboardingParams{
				Enabled: stripe.Bool(true),
			},
		},
	}
	result, err := accountsession.New(params)
	if err != nil {
		return
	}
	log.Print(result)
}
func GetAcount(ctx *gin.Context) {
	params := &stripe.AccountParams{}
	log.Print(ctx.MustGet("Stripe_Account_Id").(string))
	result, err := account.GetByID(ctx.MustGet("Stripe_Account_Id").(string), params)
	if err != nil {
		return
	}
	if !result.PayoutsEnabled {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "アカウントに口座が登録されていません。"})
	}
	ctx.JSON(http.StatusOK, gin.H{"result": result})
}
