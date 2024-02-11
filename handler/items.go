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
		itemGroup.GET("", func(ctx *gin.Context) {
			// レスポンスの処理
			PreviewList, err := items.GetPreviewList(items.ExampleItemRequests{}, ctx)
			if err != nil {
				//error logなど
				return
			}
			ctx.JSON(200, PreviewList)
		})

		itemGroup.GET("/:item_id", func(ctx *gin.Context) {

			// item_id の取得
			itemId := ctx.Param("item_id")
			if itemId == "" {
				ctx.JSON(http.StatusBadRequest, "cannot get itemId")
				return
			}
			Overview, err := items.GetOverview(items.ExampleItemRequests{}, itemId, ctx)
			if err != nil {
				return
			}
			// レスポンスの処理
			ctx.JSON(200, Overview)
		})

		itemGroup.GET("/search", func(ctx *gin.Context) {
			keywords := ctx.Query("keyword")
			log.Println(keywords)
			PreviewList, err := items.GetSearchPreviewList(items.ExampleItemRequests{}, strings.Split(keywords, "+"), ctx)
			if err != nil {
				return
			}
			ctx.JSON(200, PreviewList)
		})
	}
}
