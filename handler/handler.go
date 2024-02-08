package handler

import (
	"log"
	"strings"

	"github.com/charisworks/charisworks-backend/authstatus"
	"github.com/charisworks/charisworks-backend/items"
	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router *gin.Engine
}

func NewHandler(router *gin.Engine) *Handler {
	return &Handler{
		Router: router,
	}
}

func (h *Handler) SetupRoutes(firebaseApp *validation.FirebaseApp) {
	firebase := h.Router.Group("/firebase")
	{
		firebase.GET("/test", func(c *gin.Context) {
			idToken := "[idToken]"
			app, err := validation.NewFirebaseApp()
			if err != nil {
				return
			}
			app.VerifyIDToken(c, idToken)
		})
	}
}
func (h *Handler) SetupRoutesForItem() {
	itemGroup := h.Router.Group("/api/item")
	{
		itemGroup.GET("", func(c *gin.Context) {
			// レスポンスの処理
			PreviewList := items.GetPreviewList(items.ItemRequests{})
			c.JSON(200, PreviewList)
		})
		itemGroup.GET("/:item_id", func(c *gin.Context) {

			// item_id の取得
			itemId := c.Param("item_id")
			Overview := items.GetOverview(items.ItemRequests{}, itemId)
			// レスポンスの処理
			c.JSON(200, Overview)
		})
		itemGroup.GET("/search", func(c *gin.Context) {
			keywords := c.Query("keyword")
			log.Println(keywords)
			PreviewList := items.GetSearchPreviewList(items.ItemRequests{}, strings.Split(keywords, "+"))
			c.JSON(200, PreviewList)
		})
	}
}

func (h *Handler) SetupRoutesForAuthStatus() {
	itemGroup := h.Router.Group("/api/userauthstatus")
	{
		itemGroup.POST("", func(c *gin.Context) {
			// レスポンスの処理
			i := new(struct{ email string })
			if err := c.BindJSON(&i); err != nil {
				log.Print(err)
			}
			PreviewList := authstatus.AuthStatusCheck(i.email, authstatus.AuthStatusRequests{})
			c.JSON(200, PreviewList)
		})

	}
}
