package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func md5Hash(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// ResponseSvgWithCache send svg with needed headers
func ResponseSvgWithCache(c *gin.Context, svg string) {
	c.Header("Content-Type", "image/svg+xml;charset=utf-8")
	c.Header("Cache-Control", "no-cache")
	c.Header("ETag", md5Hash(svg))
	c.Header("Last-Modified", "Fri, 10 Aug 2018 08:42:57 GMT")
	c.String(http.StatusOK, svg)
}
