package main

import (
	"github.com/charisworks/charisworks-backend/handler"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.ContextWithFallback = true

	h := handler.NewHandler(r)

	app, err := validation.NewFirebaseApp()
	if err != nil {
		return
	}
	h.SetupRoutesForItem()
	h.SetupRoutesForAuthStatus()
	h.SetupRoutesForUser(app)

	h.Router.Run("localhost:8080")
}
