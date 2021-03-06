package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/badge-service/adapter"
	"github.com/zcong1993/badge-service/controller"
	"github.com/zcong1993/badge-service/utils"
	"os"
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
	r.GET("/bundlephobia/:topic/*rest", controller.MakeController(adapter.BundlephobiaApi, "bundlephobia", 2, "topic"))
	r.GET("/chrome-web-store/:topic/*rest", controller.MakeController(adapter.ChromeWebStoreApi, "chrome-web-store", 1, "topic"))
	r.GET("/crates/:topic/*rest", controller.MakeController(adapter.CratesApi, "crates", 1, "topic"))
	r.GET("/opencollective/:topic/*rest", controller.MakeController(adapter.OpencollectiveApi, "opencollective", 1, "topic"))
	r.GET("/packagephobia/:topic/*rest", controller.MakeController(adapter.PackagephobiaApi, "packagephobia", 2, "topic"))
	r.GET("/gem/:topic/*rest", controller.MakeController(adapter.GemApi, "gem", 2, "topic"))
	r.GET("/badge/*rest", controller.MakeController(adapter.BadgeApi, "badge", 3))

	r.Run(":" + utils.StringOrDefault(os.Getenv("PORT"), "8080"))
}
