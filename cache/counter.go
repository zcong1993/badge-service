package cache

import "github.com/zcong1993/cache/lru/counter"

var (
	store = counter.NewCounter(1000)
	// MaxParallel is max requests count for the same url
	MaxParallel = int64(5)
)

// IsBurst check if a url hit MaxParallel
func IsBurst(key interface{}) bool {
	return store.Incr(key) > MaxParallel
}

// Release decr the counter
func Release(key interface{}) {
	store.Decr(key)
}
