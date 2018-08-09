package controller

import (
	"fmt"
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
func MakeController(apiFunc adapter.ApiFunc, t string, l int, args ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		queryStyle := c.Query("style")
		style := badge.DEFAULT
		if queryStyle == "flat" {
			style = badge.FLAT
		}
		var q []string
		for _, arg := range args {
			q = append(q, c.Param(arg))
		}
		q = append(q, utils.ParamsOrDefault(c.Param("rest"), l)...)
		cacheKey := t + "-" + strings.Join(q, "-") + fmt.Sprintf("%d", style)
		cacheResp := cache.GetString(cacheKey)
		if cacheResp != "" {
			utils.ResponseSvgWithCache(c, cacheResp)
			return
		}
		res := apiFunc(q...)
		res.Style = style
		res.Subject = strings.Replace(res.Subject, "-", " ", -1)
		svg, err := badge.Badgen(badge.Input(res))
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{"ok": false})
			return
		}
		if res.Status != "api error" {
			cache.SetCache(cacheKey, string(svg), DEFULT_CACHE_AGE)
		}
		utils.ResponseSvgWithCache(c, string(svg))
	}
}
