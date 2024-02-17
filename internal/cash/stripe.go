package cash

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/account"
	"github.com/stripe/stripe-go/v76/accountlink"
	"github.com/stripe/stripe-go/v76/loginlink"
)

func CreateStripeAccount(ctx *gin.Context) error {

	email := ctx.MustGet("UserEmail").(string)
	User := ctx.MustGet("User").(*user.User)

	Account, err := GetAcount(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return err
	}
	log.Print(Account.PayoutsEnabled)
	if Account.PayoutsEnabled {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "アカウントが存在しています。"})
		return nil
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
		return err
	}
	ctx.Set("Stripe_Account_Id", a.ID)
	URL, err := CreateAccountLink(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"url": URL})
	return nil
}

func CreateAccountLink(ctx *gin.Context) (*string, error) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String(ctx.MustGet("Stripe_Account_Id").(string)),
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
func GetMypage(ctx *gin.Context) error {
	Account, err := GetAcount(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return err
	}
	if !Account.PayoutsEnabled {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "口座が登録されていません。"})
		return err
	}
	params := &stripe.LoginLinkParams{Account: stripe.String(ctx.MustGet("Stripe_Account_Id").(string))}
	result, err := loginlink.New(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"url": result.URL})
	return nil
}

func GetAcount(ctx *gin.Context) (*stripe.Account, error) {
	params := &stripe.AccountParams{}
	log.Print(ctx.MustGet("Stripe_Account_Id").(string))
	result, err := account.GetByID(ctx.MustGet("Stripe_Account_Id").(string), params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
