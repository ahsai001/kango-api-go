package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HtmlAddNews(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "addnews.html", nil)
}
