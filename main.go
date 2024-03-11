package main

import (
	"github.com/charisworks/charisworks-backend/handler"
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.ContextWithFallback = true
	utils.CORS(r)
	h := handler.NewHandler(r)
	app, err := validation.NewFirebaseApp()
	if err != nil {
		return
	}
	db, err := utils.DBInit()
	if err != nil {
		return
	}

	cartRequests := cart.Requests{CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}}
	itemRequests := items.Requests{ItemRepository: items.ItemRepository{DB: db}, ItemUtils: items.ItemUtils{}}
	userRequests := users.Requests{UserUtils: users.UserUtils{}, UserRepository: users.UserRepository{DB: db}}
	manufacturerRequests := manufacturer.Requests{ManufacturerItemRepository: manufacturer.Repository{DB: db}, ManufacturerInspectPayloadUtils: manufacturer.ManufacturerUtils{}, ItemRepository: items.ItemRepository{DB: db}}
	stripeRequests := cash.Requests{CartRequests: cartRequests, UserRequests: userRequests}
	transactionRequests := transaction.TransactionRequests{TransactionRepository: transaction.Repository{DB: db}, CartRepository: cart.Repository{DB: db}, CartUtils: cart.Utils{}, StripeRequests: cash.Requests{CartRequests: cartRequests, UserRequests: userRequests}, StripeUtils: cash.Utils{}}
	webhookRequests := transaction.Webhook{StripeUtils: cash.Utils{}, TransactionRepository: transaction.Repository{DB: db}, ItemUpdater: items.Updater{DB: db}}
	h.SetupRoutesForWebhook(webhookRequests)
	h.SetupRoutesForItem(itemRequests)
	h.SetupRoutesForUser(app, userRequests)
	h.SetupRoutesForCart(app, cartRequests, userRequests)
	h.SetupRoutesForManufacturer(app, manufacturerRequests)
	h.SetupRoutesForStripe(app, userRequests, stripeRequests, transactionRequests)
	h.Router.Run("localhost:8080")
}
