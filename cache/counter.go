package cache

import "github.com/zcong1993/cache/lru/counter"

var (
	store       = counter.NewCounter(1000)
	MaxParallel = int64(5)
)

func IsBurst(key interface{}) bool {
	return store.Incr(key) > MaxParallel
}

func Release(key interface{}) {
	store.Decr(key)
}
