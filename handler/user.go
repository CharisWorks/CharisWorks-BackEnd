package handler

import (
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForUser(firebaseApp *validation.FirebaseApp) {
	UserRouter := h.Router.Group("/api")
	UserRouter.Use(FirebaseMiddleware(*firebaseApp))
	{
		UserRouter.GET("/user", func(c *gin.Context) {
			idToken := "[idToken]"
			app, err := validation.NewFirebaseApp()
			if err != nil {
				return
			}
			app.VerifyIDToken(c, idToken)
		})
	}
}
