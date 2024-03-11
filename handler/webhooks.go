package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func (h *Handler) SetupRoutesForWebhook(webhookRequests transaction.IWebhook) {
	UserRouter := h.Router.Group("/webhook")
	UserRouter.Use(webhookMiddleware())
	{
		UserRouter.POST("", func(ctx *gin.Context) {
			event := ctx.MustGet("event").(stripe.Event)
			// Handle the checkout.session.completed event
			if event.Type == "checkout.session.completed" {
				var sessions stripe.CheckoutSession
				err := json.Unmarshal(event.Data.Raw, &sessions)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
					return
				}

				params := &stripe.CheckoutSessionParams{}
				params.AddExpand("line_items")
				log.Print(sessions.ID)

				err = webhookRequests.PurchaseComplete(sessions.ID)
				if err != nil {
					return
				}

			}
			if event.Type == "checkout.session.cancelled" {
				var sessions stripe.CheckoutSession
				json.Unmarshal(event.Data.Raw, &sessions)
				log.Print(sessions.ID)

				webhookRequests.PurchaseCanceled(sessions.ID)

			}
			if event.Type == "checkout.session.failed" {
				var sessions stripe.CheckoutSession
				json.Unmarshal(event.Data.Raw, &sessions)
				log.Print(sessions.ID)
				webhookRequests.PurchaseFail(sessions.ID)

			}
		})
	}
}
