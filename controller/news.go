package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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
	result, err := repo.GetNewsDetail(id)
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

}
