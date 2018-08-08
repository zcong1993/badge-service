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

	r.GET("/circleci/*rest", controller.MakeController(adapter.CircleciApi, "circleci", 4))
	r.GET("/docker/:topic/*rest", controller.MakeController(adapter.DockerApi, "docker", 2, "topic"))
	r.GET("/github/:topic/*rest", controller.MakeController(adapter.GithubApi, "github", 2, "topic"))
	r.GET("/travis/*rest", controller.MakeController(adapter.TravisApi, "travis", 3))
	r.GET("/npm/:topic/*rest", controller.MakeController(adapter.NpmApi, "npm", 3, "topic"))
	r.GET("/homebrew/:topic/*rest", controller.MakeController(adapter.HomebrewApi, "homebrew", 1, "topic"))
	r.GET("/pypi/:topic/*rest", controller.MakeController(adapter.PypiApi, "pypi", 1, "topic"))
	r.GET("/codecov/*rest", controller.MakeController(adapter.CodecovApi, "codecov", 4))
	r.GET("/appveyor/*rest", controller.MakeController(adapter.AppveyorApi, "appveyor", 3))

	r.Run()
}
