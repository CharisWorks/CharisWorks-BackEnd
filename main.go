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

	app := validation.NewFirebaseApp()

	h.SetupRoutes(app)

	h.Router.Run("localhost:8080")
}
