package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseSvgWithCache send svg with needed headers
func ResponseSvgWithCache(c *gin.Context, svg string) {
	c.Header("Content-Type", "image/svg+xml;charset=utf-8")
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.String(http.StatusOK, string(svg))
}
