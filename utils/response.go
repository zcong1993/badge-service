package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"crypto/md5"
)

func md5Hash(str string) string {
	data := []byte(str)
	has := md5.Sum(data)

	return fmt.Sprintf("%x", has)
}

// ResponseSvgWithCache send svg with needed headers
func ResponseSvgWithCache(c *gin.Context, svg string) {
	c.Header("Content-Type", "image/svg+xml;charset=utf-8")
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Header("ETag", md5Hash(svg))
	c.String(http.StatusOK, svg)
}
