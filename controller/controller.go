package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/badge"
	"github.com/zcong1993/badge-service/adapter"
	"github.com/zcong1993/badge-service/cache"
	"github.com/zcong1993/badge-service/utils"
	"net/http"
	"strings"
	"time"
)

// DEFULT_CACHE_AGE is expire time of cache
const DEFULT_CACHE_AGE = time.Second * 30

// MakeController return a gin handler function
func MakeController(apiFunc adapter.ApiFunc, l int, args ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		queryStyle := c.Query("style")
		var q []string
		for _, arg := range args {
			q = append(q, c.Param(arg))
		}
		q = append(q, utils.ParamsOrDefault(c.Param("rest"), l)...)
		cacheKey := strings.Join(q, "-") + "-" + queryStyle
		cacheResp := cache.GetString(cacheKey)
		if cacheResp != "" {
			println("hint cache")
			utils.ResponseSvgWithCache(c, cacheResp)
			return
		}
		res := apiFunc(q...)
		style := badge.DEFAULT
		if queryStyle == "flat" {
			style = badge.FLAT
		}
		res.Style = style
		svg, err := badge.Badgen(badge.Input(res))
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{"ok": false})
			return
		}
		cache.SetCache(cacheKey, string(svg), DEFULT_CACHE_AGE)
		utils.ResponseSvgWithCache(c, string(svg))
	}
}
