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
	h.SetupRoutesForItem(items.ExampleItemRequests{}, items.ExampleItemDB{}, items.ExampleItemUtils{})
	h.SetupRoutesForUser(app, users.UserRequests{}, users.UserDB{DB: db})
	h.SetupRoutesForCart(app, cart.CartRequests{}, cart.ExampleCartDB{}, users.UserRequests{}, cart.CartUtils{}, users.UserDB{})
	h.SetupRoutesForManufacturer(app, manufacturer.ExampleManufacturerRequests{})
	h.SetupRoutesForStripe(app, cash.ExampleTransactionRequests{}, cash.StripeRequests{}, cart.CartRequests{}, cart.ExampleCartDB{}, cart.CartUtils{}, items.ExampleItemDB{}, cash.ExampleTransactionDBHistory{}, users.UserRequests{}, users.UserDB{DB: db})
	h.Router.Run("localhost:8080")
}
