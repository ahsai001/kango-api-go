package main

import (
	"fmt"
	"kango-api/controller"
	"kango-api/helper/filehelper"

	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	filehelper.InitFolder("upload/photo/")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	//setting
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	//define some routes
	r.GET("/", controller.Welcome)
	r.POST("/cijou/news/all", controller.GetNews)
	r.POST("/cijou/news/detail/:id", controller.GetNewsDetail)
	r.POST("/cijou/news/add", controller.AddNews)
	r.POST("/cijou/login", controller.Login)

	r.Run(fmt.Sprintf(":%s", port))
}
