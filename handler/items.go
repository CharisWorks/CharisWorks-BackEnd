package handler

import (
	"strings"

	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForItem(itemsImpl items.IItemRequests) {
	itemGroup := h.Router.Group("/api/item")
	{
		itemGroup.GET("", func(ctx *gin.Context) {
			// レスポンスの処理
			PreviewList, err := itemsImpl.GetPreviewList(ctx)
			if err != nil {
				//error logなど
				return
			}
			ctx.JSON(200, PreviewList)
		})

		itemGroup.GET("/:item_id", func(ctx *gin.Context) {

			// item_id の取得
			itemId, err := getParams("item_id", ctx)
			if err != nil {
				return
			}
			Overview, err := itemsImpl.GetOverview(*itemId, ctx)
			if err != nil {
				return
			}
			// レスポンスの処理
			ctx.JSON(200, Overview)
		})

		itemGroup.GET("/search", func(ctx *gin.Context) {
			keywords, err := getQuery("keyword", ctx)
			if err != nil {
				return
			}
			PreviewList, err := itemsImpl.GetSearchPreviewList(strings.Split(*keywords, "+"), ctx)
			if err != nil {
				return
			}
			ctx.JSON(200, PreviewList)
		})
	}
}
