package cash

import (
	"log"
	"net/http"

	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func CreatePaymentIntent(ctx *gin.Context, u ITransactionUtils, c cart.ICartRequest) (*string, error) {
	Carts, err := c.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = u.InspectCart(*Carts)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}
	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(u.GetTotalAmount(*Carts)), //合計金額を算出する関数をインジェクト
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
	ctx.JSON(http.StatusOK, gin.H{"clientSecret": pi.ClientSecret})
	return &pi.ClientSecret, nil
}
func GetTransactionList(ctx *gin.Context, i ITransactionRequests) error {
	TransactionList, err := i.GetTransactionList(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, TransactionList)
	return nil
}
func GetTransactionDetails(TransactionId string, ctx *gin.Context, i ITransactionRequests) error {
	TransactionList, err := i.GetTransactionDetails(ctx, TransactionId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, TransactionList)
	return nil
}
