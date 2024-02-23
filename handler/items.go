package handler

import (
	"strings"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForItem(ItemRequests items.IItemRequests, ItemDB items.IItemDB, ItemUtils items.IItemUtils) {
	itemGroup := h.Router.Group("/api/item")
	{
		itemGroup.GET("/:item_id", func(ctx *gin.Context) {

			// item_id の取得
			itemId, err := getParams("item_id", true, ctx)
			if err != nil {
				return
			}
			Overview, err := ItemRequests.GetOverview(*itemId, ItemDB, ctx)
			if err != nil {
				return
			}
			// レスポンスの処理
			ctx.JSON(200, Overview)
		})

		itemGroup.GET("/", func(ctx *gin.Context) {
			page, _ := getQuery("page", false, ctx)
			sort, _ := getQuery("sort", false, ctx)
			keywords, _ := getQuery("keyword", false, ctx)
			keywordlist := strings.Split(*keywords, "+")
			manufacturer, _ := getQuery("manufacturer", false, ctx)
			PreviewList, err := ItemRequests.GetSearchPreviewList(&keywordlist, page, sort, manufacturer, ItemDB, ItemUtils, ctx)
			if err != nil {
				return
			}
			ctx.JSON(200, PreviewList)
		})
	}
}
