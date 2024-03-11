package cash

import (
	"log"
	"regexp"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/account"
	"github.com/stripe/stripe-go/v76/accountlink"
	"github.com/stripe/stripe-go/v76/loginlink"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/ttacon/libphonenumber"
)

type Requests struct {
	CartRequests cart.IRequests
	UserRequests users.IRequests
}

func (r Requests) GetRegisterLink(email string, user users.User) (url string, err error) {
	log.Print(email)
	Account, err := GetAccount(user.UserProfile.StripeAccountId)
	if err != nil {
		return url, &utils.InternalError{Message: utils.InternalErrorNotFound}
	}
	if Account.PayoutsEnabled {
		return url, &utils.InternalError{Message: utils.InternalErrorManufacturerAlreadyHasBank}
	}
	if &user.UserAddress == new(users.UserAddress) {
		return url, &utils.InternalError{Message: utils.InternalErrorAccountIsNotSatisfied}
	}
	pnum, err := libphonenumber.Parse(user.UserAddress.PhoneNumber, "JP")
	e164Number := new(string)
	if err != nil {
		return url, &utils.InternalError{Message: utils.InternalErrorIncident}
	}
	*e164Number = libphonenumber.Format(pnum, libphonenumber.E164)
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
			FirstNameKanji: stripe.String(user.UserAddress.FirstName),
			FirstNameKana:  stripe.String(user.UserAddress.FirstNameKana),
			LastNameKanji:  stripe.String(user.UserAddress.LastName),
			LastNameKana:   stripe.String(user.UserAddress.LastNameKana),
			Email:          stripe.String(email),
			Phone:          stripe.String(*e164Number),
		},
		BusinessProfile: &stripe.AccountBusinessProfileParams{
			MCC:                stripe.String("5699"),
			URL:                stripe.String("charis.works/user/profile/" + user.UserId),
			ProductDescription: stripe.String("this is an account of manufacturer for charis works"),
		},
	}

	a, err := account.New(params)
	if err != nil {
		log.Print("Stripe Error: ", err)
		return url, &utils.InternalError{Message: utils.InternalErrorFromStripe}
	}
	err = r.UserRequests.ProfileUpdate(user.UserId, users.UserProfile{StripeAccountId: a.ID})
	if err != nil {
		return url, err
	}
	accountLinkParams := &stripe.AccountLinkParams{
		Account:    stripe.String(a.ID),
		RefreshURL: stripe.String("http://localhost:3000"),
		ReturnURL:  stripe.String("http://localhost:3000"),
		Type:       stripe.String("account_onboarding"),
		Collect:    stripe.String("eventually_due"),
	}
	accountLink, err := accountlink.New(accountLinkParams)
	if err != nil {
		log.Print("Stripe Error: ", err)
		return url, &utils.InternalError{Message: utils.InternalErrorFromStripe}
	}
	return accountLink.URL, nil

}

func (r Requests) GetStripeMypageLink(stripeAccountId string) (url string, err error) {
	Account, err := GetAccount(stripeAccountId)
	if err != nil {
		return url, err
	}
	if !Account.PayoutsEnabled {
		return url, &utils.InternalError{Message: utils.InternalErrorManufacturerDoesNotHaveBank}
	}
	params := &stripe.LoginLinkParams{Account: &Account.ID}
	result, err := loginlink.New(params)
	if err != nil {
		log.Print("Stripe Error: ", err)
		return url, &utils.InternalError{Message: utils.InternalErrorFromStripe}
	}
	return result.URL, nil
}

func GetAccount(stripeAccountId string) (*stripe.Account, error) {
	params := &stripe.AccountParams{}
	regex := regexp.MustCompile(`acct_\w+`)
	matches := regex.FindAllString(stripeAccountId, -1)
	for _, match := range matches {
		if regex.MatchString(match) {
			result, err := account.GetByID(stripeAccountId, params)
			if err != nil {
				log.Print("Stripe Error: ", err)
				return nil, &utils.InternalError{Message: utils.InternalErrorFromStripe}
			}
			return result, nil
		}
	}
	result := new(stripe.Account)
	result.PayoutsEnabled = false
	log.Print(result)
	return result, nil

}

func (r Requests) CreatePaymentintent(userId string, totalAmount int) (ClientSecret string, StripeTransactionId string, err error) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthGP4F3QjdR0SKk77E4pGHrsBAQEHia6lasXyujFOKXDyrodAxaE6PH6u2kNCVSdC5dBIRh82u00XqHQIZjM"

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(totalAmount)), //合計金額を算出する関数をインジェクト

		Currency: stripe.String(string(stripe.CurrencyJPY)),
		// In the latest version of the API, specifying the `automatic_payment_methods` parameter is optional because Stripe enables its functionality by default.
		PaymentMethodTypes: []*string{stripe.String("card"), stripe.String("konbini")},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)
	if err != nil {
		log.Printf("pi.New: %v", err)
		log.Print("Stripe Error: ", err)
		return ClientSecret, StripeTransactionId, &utils.InternalError{Message: utils.InternalErrorFromStripe}
	}
	return pi.ClientSecret, pi.ID, nil
}
