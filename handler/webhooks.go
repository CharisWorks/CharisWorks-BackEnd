package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charisworks/charisworks-backend/internal/admin"
	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func (h *Handler) SetupRoutesForWebhook(webhookRequests transaction.IWebhook, app validation.IFirebaseApp) {
	UserRouter := h.Router.Group("/webhook")
	UserRouter.Use(webhookMiddleware())
	{
		UserRouter.POST("", func(ctx *gin.Context) {
			event := ctx.MustGet(string(stripeEvent)).(stripe.Event)

			if event.Type == "payment_intent.succeeded" {
				var sessions stripe.CheckoutSession
				err := json.Unmarshal(event.Data.Raw, &sessions)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
					return
				}

				params := &stripe.CheckoutSessionParams{}
				params.AddExpand("line_items")
				log.Print(sessions.ID)

				transactionDetails, err := webhookRequests.PurchaseComplete(sessions.ID)
				if err != nil {
					return
				}
				admin.SendPurchasedEmail(transactionDetails, app)

			}
		})
	}
}
