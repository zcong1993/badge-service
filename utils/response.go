package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseSvgWithCache send svg with needed headers
func ResponseSvgWithCache(c *gin.Context, svg string) {
	c.Header("Content-Type", "image/svg+xml;charset=utf-8")
	c.Header("Cache-Control", "public, max-age=60, stale-while-revalidate=86400, stale-if-error=86400, s-maxage=86400")
	c.String(http.StatusOK, string(svg))
}
