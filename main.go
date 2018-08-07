package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/badge-service/controller"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/circleci/*rest", controller.CircleciController)
	r.GET("/docker/:topic/*rest", controller.DockerController)

	r.Run()
}
