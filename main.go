package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/badge-service/adapter"
	"github.com/zcong1993/badge-service/controller"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/circleci/*rest", controller.MakeController(adapter.CircleciApi, 4))
	r.GET("/docker/:topic/*rest", controller.MakeController(adapter.DockerApi, 2, "topic"))
	r.GET("/github/:topic/*rest", controller.MakeController(adapter.GithubApi, 2, "topic"))
	r.GET("/travis/*rest", controller.MakeController(adapter.TravisApi, 3))

	r.Run()
}
