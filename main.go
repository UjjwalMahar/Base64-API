package main

import (
	"github.com/UjjwalMahar/base64-api/controllers"
	"github.com/UjjwalMahar/base64-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.GetRoot)
	r.POST("/encode", controllers.PostEncode)
	r.POST("/decode", controllers.PostDecode)
	r.Run() 
}