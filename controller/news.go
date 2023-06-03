package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"kango-api/helper/filehelper"
	models "kango-api/model"
	repo "kango-api/repository"
)

func GetNews(ctx *gin.Context) {
	result, err := repo.GetNews()
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewsResponse{
			Status:  false,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func GetNewsDetail(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	result, err := repo.GetNewsDetail(int64(id))
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewsDetailResponse{
			Status:  false,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, models.NewsDetailResponse{
			Status:  true,
			Message: "Sukses",
			Data:    *result,
		})
	}
}

func AddNews(ctx *gin.Context) {
	filePath, err := filehelper.UploadSingleFile(ctx.Request, "photo", "upload/photo/")
	if err != nil {
		log.Println(err)
	}
	title := ctx.PostForm("title")
	summary := ctx.PostForm("summary")
	body := ctx.PostForm("body")

	detail, err := repo.AddNews(title, summary, body, filePath)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewsDetailResponse{
			Status:  false,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, models.NewsDetailResponse{
			Status:  true,
			Message: "Sukses",
			Data:    *detail,
		})
	}
}
