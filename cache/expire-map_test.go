package cache

import (
	"testing"
	"time"
)

var e = NewExpireMap(time.Minute)

func BenchmarkExpireMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e.Set("hello", "world", time.Second)
		e.Get("hello")
	}
}
