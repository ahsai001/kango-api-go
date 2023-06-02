package main

import (
	"kango-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//define some routes
	r.GET("/", func(ctx *gin.Context) {
		controller.Welcome(ctx)
	})
	r.GET("/cijou/news/all", controller.GetNews)
	r.GET("/cijou/news/detail/:id", controller.GetNewsDetail)
	r.POST("/cijou/login", controller.Login)
	r.POST("/cijou/news/add", controller.AddNews)

	r.Run(":8080")
}
