package handler

import (
	"log"
	"strings"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForItem() {
	itemGroup := h.Router.Group("/api/item")
	{
		itemGroup.GET("", func(c *gin.Context) {
			// レスポンスの処理
			PreviewList := items.GetPreviewList(items.ExampleItemRequests{})
			c.JSON(200, PreviewList)
		})

		itemGroup.GET("/:item_id", func(c *gin.Context) {

			// item_id の取得
			itemId := c.Param("item_id")
			Overview := items.GetOverview(items.ExampleItemRequests{}, itemId)
			// レスポンスの処理
			c.JSON(200, Overview)
		})

		itemGroup.GET("/search", func(c *gin.Context) {
			keywords := c.Query("keyword")
			log.Println(keywords)
			PreviewList := items.GetSearchPreviewList(items.ExampleItemRequests{}, strings.Split(keywords, "+"))
			c.JSON(200, PreviewList)
		})
	}
}
