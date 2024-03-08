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

func (r Utils) Refund(amount int, transactionId string, transferId string) (err error) {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthx8UK5v9CdaucJOZj3FwkOHZ8KjDt25IAvplosSab4uybQOyE2Ne6xxxI4Rnh8pWEbYUwPoPG00wvseAHzl"

	params := &stripe.RefundParams{PaymentIntent: stripe.String(transactionId)}
	result, err := refund.New(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
	stripe.Key = "sk_test_51Nj1urA3bJzqElthx8UK5v9CdaucJOZj3FwkOHZ8KjDt25IAvplosSab4uybQOyE2Ne6xxxI4Rnh8pWEbYUwPoPG00wvseAHzl"
	log.Print("Reversing transfer... \n amount: ", amount, "\n transferId: ", transferId)

	reverseParams := &stripe.TransferReversalParams{
		Amount: stripe.Int64(int64(amount)),
		ID:     stripe.String(transferId),
	}
	transferResult, err := transferreversal.New(reverseParams)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(transferResult)
	return
}
func (r Utils) Transfer(amount int, stripeAccountId string, transactionId string) *string {
	stripe.Key = "sk_test_51Nj1urA3bJzqElthx8UK5v9CdaucJOZj3FwkOHZ8KjDt25IAvplosSab4uybQOyE2Ne6xxxI4Rnh8pWEbYUwPoPG00wvseAHzl"
	log.Print("Transfering... \n amount: ", amount, "\n stripeID: ", stripeAccountId, "\n transactionId: ", transactionId)
	params := &stripe.TransferParams{
		Amount:      stripe.Int64(int64(amount)),
		Currency:    stripe.String(string(stripe.CurrencyJPY)),
		Destination: stripe.String(stripeAccountId),
		Description: stripe.String(transactionId),
	}
	tr, _ := transfer.New(params)
	log.Print(tr.ID)
	return &tr.ID
}
