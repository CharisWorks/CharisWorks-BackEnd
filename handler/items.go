package handler

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForItem(ItemRequests items.IItemRequests, ItemDB items.IItemDB, ItemUtils items.IItemUtils) {
	itemGroup := h.Router.Group("/api/item")
	{
		itemGroup.GET("/:item_id", func(ctx *gin.Context) {
			itemId, err := utils.GetParams("item_id", ctx)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			Overview, err := ItemRequests.GetOverview(*itemId, ItemDB)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			// レスポンスの処理
			ctx.JSON(200, Overview)
		})

		itemGroup.GET("/", func(ctx *gin.Context) {

			PreviewList, err := ItemRequests.GetSearchPreviewList(ctx, ItemDB, ItemUtils)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(200, PreviewList)
		})
	}
}
