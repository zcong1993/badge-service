package cache

import (
	"github.com/zcong1993/cache/expire"
	"time"
)

var em = expire.NewStringExpireMap(time.Hour)

// GetString get string cache by key
func GetString(key string) string {
	str := em.Get(key)
	if str == nil {
		return ""
	}
	return *str
}

// SetCache add cache key value
func SetCache(k string, val string, expire time.Duration) {
	em.Set(k, val, expire)
}
