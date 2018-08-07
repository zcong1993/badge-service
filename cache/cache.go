package cache

import "time"

var em = NewExpireMap(time.Minute)

func GetString(key string) string {
	str, ok := em.Get(key).(string)
	if !ok {
		return ""
	}
	return str
}

func SetCache(k string, val interface{}, expire time.Duration) {
	em.Set(k, val, expire)
}
