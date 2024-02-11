package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForItem() {
	itemGroup := h.Router.Group("/api/item")
	{
		itemGroup.GET("", func(c *gin.Context) {
			// レスポンスの処理
			PreviewList, err := items.GetPreviewList(c, items.ExampleItemRequests{})
			if err != nil {
				//error logなど
				return
			}
			c.JSON(200, PreviewList)
		})

		itemGroup.GET("/:item_id", func(c *gin.Context) {

			// item_id の取得
			itemId := c.Param("item_id")
			if itemId == "" {
				c.JSON(http.StatusBadRequest, "cannot get itemId")
				return
			}
			Overview, err := items.GetOverview(items.ExampleItemRequests{}, itemId)
			if err != nil {
				c.JSON(http.StatusNotFound, err)
			}
			// レスポンスの処理
			c.JSON(200, Overview)
		})

		itemGroup.GET("/search", func(c *gin.Context) {
			keywords := c.Query("keyword")
			log.Println(keywords)
			PreviewList, err := items.GetSearchPreviewList(items.ExampleItemRequests{}, strings.Split(keywords, "+"))
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.JSON(200, PreviewList)
		})
	}
}
