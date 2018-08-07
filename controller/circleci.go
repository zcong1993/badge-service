package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/badge"
	"github.com/zcong1993/badge-service/adapter"
	"github.com/zcong1993/badge-service/cache"
	"github.com/zcong1993/badge-service/utils"
	"net/http"
)

func CircleciController(c *gin.Context) {
	queryStyle := c.Query("style")
	rest := c.Param("rest")
	cacheKey := queryStyle + rest
	cacheResp := cache.GetString(cacheKey)
	if cacheResp != "" {
		println("hint cache")
		utils.ResponseSvgWithCache(c, cacheResp)
		return
	}
	p := utils.ParamsOrDefault(rest, 4)
	res := adapter.CircleciApi(p[0], p[1], p[2], p[3])
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
