package handler

import (
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForItem(itemRequests items.IRequests) {
	itemGroup := h.Router.Group("/api/item")
	{
		itemGroup.GET("/:item_id", func(ctx *gin.Context) {
			itemId, err := utils.GetParams("item_id", ctx)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			Overview, err := itemRequests.GetOverview(*itemId)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			// レスポンスの処理
			ctx.JSON(200, Overview)
		})

		itemGroup.GET("", func(ctx *gin.Context) {
			PreviewList, totalElements, err := itemRequests.GetSearchPreviewList(ctx)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"previewList": PreviewList, "totalElements": totalElements})
		})
	}
}
