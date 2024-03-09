package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func (h *Handler) SetupRoutesForWebhook() {
	UserRouter := h.Router.Group("/webhook")
	UserRouter.Use(webhookMiddleware())
	{
		UserRouter.POST("", func(ctx *gin.Context) {

		})
	}
}

func SetupRoutesForWebhook(event stripe.Event, webhookRequests transaction.IWebhook) error {

	// Handle the checkout.session.completed event
	if event.Type == "checkout.session.completed" {
		var sessions stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &sessions)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)

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
