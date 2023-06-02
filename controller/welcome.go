package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Halo kango-api",
	})
}
