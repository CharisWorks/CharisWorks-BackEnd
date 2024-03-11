package cash

import (
	"log"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/refund"
	"github.com/stripe/stripe-go/v76/transfer"
	"github.com/stripe/stripe-go/v76/transferreversal"
)

type Utils struct {
}

func (r Utils) Refund(amount int, transferId string, accountId string) (err error) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthGP4F3QjdR0SKk77E4pGHrsBAQEHia6lasXyujFOKXDyrodAxaE6PH6u2kNCVSdC5dBIRh82u00XqHQIZjM"

	params := &stripe.RefundParams{PaymentIntent: stripe.String(transferId)}
	result, err := refund.New(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
	stripe.Key = "sk_test_51Nj1urA3bJzqElthGP4F3QjdR0SKk77E4pGHrsBAQEHia6lasXyujFOKXDyrodAxaE6PH6u2kNCVSdC5dBIRh82u00XqHQIZjM"
	log.Print("Reversing transfer... \n amount: ", amount, "\n transferId: ", accountId)

	reverseParams := &stripe.TransferReversalParams{
		Amount: stripe.Int64(int64(amount)),
		ID:     stripe.String(accountId),
	}
	transferResult, err := transferreversal.New(reverseParams)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(transferResult)
	return
}
func (r Utils) Transfer(amount int, stripeAccountId string, transactionId string) *string {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthGP4F3QjdR0SKk77E4pGHrsBAQEHia6lasXyujFOKXDyrodAxaE6PH6u2kNCVSdC5dBIRh82u00XqHQIZjM"
	log.Print("Transfering... \n amount: ", amount, "\n stripeID: ", stripeAccountId, "\n transactionId: ", transactionId)
	params := &stripe.TransferParams{
		Amount:      stripe.Int64(int64(amount)),
		Currency:    stripe.String(string(stripe.CurrencyJPY)),
		Destination: stripe.String(stripeAccountId),
		Description: stripe.String(transactionId),
	}
	tr, err := transfer.New(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(tr.ID)
	return &tr.ID
}
