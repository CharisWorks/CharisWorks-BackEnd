package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/charisworks/charisworks-backend/internal/images"
	"github.com/charisworks/charisworks-backend/internal/items"
	"github.com/charisworks/charisworks-backend/internal/manufacturer"
	"github.com/charisworks/charisworks-backend/internal/users"
	"github.com/charisworks/charisworks-backend/internal/utils"

	"github.com/charisworks/charisworks-backend/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutesForImages(firebaseApp validation.IFirebaseApp, manufacturerRequests manufacturer.IItemRequests, itemRequests items.IRequests, userRequests users.IRequests) {
	Crud := images.R2Conns{Crud: nil}
	Crud.Init()
	UserRouter := h.Router.Group("/images")
	{
		UserRouter.GET("/:item_id", func(ctx *gin.Context) {
			itemId, err := utils.GetParams("item_id", ctx)
			if err != nil {
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			log.Print(*itemId)
			items, err := Crud.GetImages(*itemId + "/")
			if err != nil {
				log.Print(err)
				err = &utils.InternalError{Message: utils.InternalErrorR2}
				utils.ReturnErrorResponse(ctx, err)
				return
			}
			log.Print(items)
			if len(items) == 0 {
				utils.ReturnErrorResponse(ctx, &utils.InternalError{Message: utils.InternalErrorNotFound})
				return
			}
			for i, item := range items {
				item = os.Getenv("IMAGES_URL") + "/" + item
				items[i] = item
			}
			ctx.JSON(http.StatusOK, items)
		})
	}
	UserRouter.Use(firebaseMiddleware(firebaseApp))
	{
		UserRouter.Use((userMiddleware(userRequests)))
		UserRouter.Use(manufacturerMiddleware())
		{
			UserRouter.POST("/:item_id", func(ctx *gin.Context) {
				itemId, err := utils.GetParams("item_id", ctx)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				userId := ctx.GetString("userId")
				itemOverview, err := itemRequests.GetOverview(*itemId)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				if itemOverview.Manufacturer.UserId != userId {
					utils.ReturnErrorResponse(ctx, &utils.InternalError{Message: utils.InternalErrorUnAuthorized})
					return
				}
				images, err := Crud.GetImages(*itemId + "/")
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				for _, image := range images {
					Crud.DeleteImage(image)
				}
				ctx.Request.ParseMultipartForm(32 << 20)
				files := ctx.Request.MultipartForm.File["photo"]
				if len(files) == 0 {
					ctx.String(400, "Bad request")
					return
				}

				// ファイルを保存
				for i, file := range files {
					// ファイルのオープン
					destPath := ""
					src, err := file.Open()
					if err != nil {
						ctx.String(http.StatusInternalServerError, "ファイルを開けませんでした")
						return
					}
					defer src.Close()
					b, err := io.ReadAll(src)
					if err != nil {
						ctx.String(http.StatusInternalServerError, "ファイルを開けませんでした")
						return
					}
					// ファイルの保存先パス
					if i == 0 {
						destPath = filepath.Join(*itemId, "/thumb.png")

					} else {
						destPath = filepath.Join(*itemId, "/", strconv.Itoa(i)+".png")
					}
					Crud.Crud.UploadObject(ctx, b, destPath)
				}

				ctx.String(200, "File uploaded")
			})
			UserRouter.DELETE("/:item_id", func(ctx *gin.Context) {
				itemId, err := utils.GetParams("item_id", ctx)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				userId := ctx.GetString("userId")
				itemOverview, err := itemRequests.GetOverview(*itemId)
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				if itemOverview.Manufacturer.UserId != userId {
					utils.ReturnErrorResponse(ctx, &utils.InternalError{Message: utils.InternalErrorUnAuthorized})
					return
				}
				images, err := Crud.GetImages(*itemId + "/")
				if err != nil {
					utils.ReturnErrorResponse(ctx, err)
					return
				}
				for _, image := range images {
					Crud.DeleteImage(image)
				}
			})
		}

	}
}
