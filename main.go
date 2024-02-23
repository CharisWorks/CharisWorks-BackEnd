package main

import (
	"github.com/charisworks/charisworks-backend/handler"
	"github.com/charisworks/charisworks-backend/internal/cart"
	"github.com/charisworks/charisworks-backend/internal/cash"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/user"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.ContextWithFallback = true
	handler.CORS(r)
	h := handler.NewHandler(r)

	app, err := validation.NewFirebaseApp()
	if err != nil {
		return
	}
	h.SetupRoutesForItem(items.ExampleItemRequests{}, items.ExampleItemDB{}, items.ExampleItemUtils{})
	h.SetupRoutesForUser(app, user.ExampleUserRequests{})
	h.SetupRoutesForCart(app, cart.CartRequests{}, cart.ExampleCartDB{}, user.ExampleUserRequests{}, cart.CartUtils{})
	h.SetupRoutesForManufacturer(app, manufacturer.ExampleManufacturerRequests{})
	h.SetupRoutesForStripe(app, cash.ExampleTransactionRequests{}, cash.StripeRequests{}, cart.CartRequests{}, cart.ExampleCartDB{}, cart.CartUtils{}, items.ExampleItemDB{})
	h.Router.Run("localhost:8080")
}
