package main

import (
	"github.com/charisworks/charisworks-backend/handler"
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
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
	manufacturerRequests := manufacturer.Requests{ManufacturerItemRepository: manufacturer.Repository{DB: db}, ManufacturerInspectPayloadUtils: manufacturer.ManufacturerUtils{}}
	h.SetupRoutesForItem(itemRequests)
	h.SetupRoutesForUser(app, userRequests)
	h.SetupRoutesForCart(app, cartRequests, userRequests)
	h.SetupRoutesForManufacturer(app, manufacturerRequests)
	h.SetupRoutesForStripe(app, cash.ExampleTransactionRequests{}, cash.StripeRequests{}, cartRequests, cart.Repository{}, cart.Utils{}, items.ItemRepository{}, cash.ExampleTransactionDBHistory{}, users.Requests{}, users.UserRepository{DB: db})
	h.Router.Run("localhost:8080")
}
