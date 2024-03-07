package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
)

func PaymentComplete(ctx *gin.Context, webhookRequests transaction.IWebhook) error {
	const MaxBodyBytes = int64(65536)
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, MaxBodyBytes)

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
		return &utils.InternalError{Message: utils.InternalErrorFromStripe}
	}

	// Pass the request body and Stripe-Signature header to ConstructEvent, along with the webhook signing key
	// You can find your endpoint's secret in your webhook settings
	endpointSecret := os.Getenv("STRIPE_KEY")
	event, err := webhook.ConstructEvent(body, ctx.Request.Header.Get("Stripe-Signature"), endpointSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest) // Return a 400 error on a bad signature

		return &utils.InternalError{Message: utils.InternalErrorFromStripe}
	}

	// Handle the checkout.session.completed event
	if event.Type == "checkout.session.completed" {
		var sessions stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &sessions)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return &utils.InternalError{Message: utils.InternalErrorFromStripe}
		}

		params := &stripe.CheckoutSessionParams{}
		params.AddExpand("line_items")
		log.Print(sessions.ID)
		err = webhookRequests.PurchaseComplete(sessions.ID)
		if err != nil {
			return err
		}
		return nil
	}
	if event.Type == "checkout.session.cancelled" {
		var sessions stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &sessions)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return &utils.InternalError{Message: utils.InternalErrorFromStripe}
		}
		log.Print(sessions.ID)
		err = webhookRequests.PurchaseCanceled(sessions.ID)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
