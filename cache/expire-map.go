package cache

import (
	"errors"
	"sync"
	"time"
)

var (
	// NO_KEY_TO_UPDATE is error message
	NO_KEY_TO_UPDATE = errors.New("key not exists or expired, set it first. ")
)

type ExipreMap struct {
	store      map[string]Value
	GcInterval time.Duration
	inGc       bool
	mu         *sync.RWMutex
	t          *time.Ticker
}

type Value struct {
	Val       interface{}
	expiredIn time.Time
}

func NewExpireMap(interval time.Duration) *ExipreMap {
	em := &ExipreMap{
		store:      make(map[string]Value),
		GcInterval: interval,
		inGc:       false,
		mu:         new(sync.RWMutex),
	}
	go em.startGc()
	return em
}

func (em *ExipreMap) startGc() {
	t := time.NewTicker(em.GcInterval)
	em.t = t
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if em.inGc {
				return
			}
			em.Gc()
		}
	}
}

func (em *ExipreMap) isExpired(k string) bool {
	em.mu.RLock()
	defer em.mu.RUnlock()
	v, ok := em.store[k]
	if !ok {
		return false
	}
	return isExpiredValue(v)
}

func isExpiredValue(val Value) bool {
	if val.expiredIn.UnixNano() <= time.Now().UnixNano() {
		return true
	}
	return false
}

func (em *ExipreMap) Gc() {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.inGc = true
	for k, v := range em.store {
		if isExpiredValue(v) {
			delete(em.store, k)
		}
	}
	em.inGc = false
}

func (em *ExipreMap) Get(key string) interface{} {
	em.mu.RLock()
	defer em.mu.RUnlock()
	v, ok := em.store[key]
	if !ok {
		return nil
	}
	if isExpiredValue(v) {
		return nil
	}
	return v.Val
}

func (em *ExipreMap) Set(k string, val interface{}, expire time.Duration) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.store[k] = Value{
		Val:       val,
		expiredIn: time.Now().Add(expire),
	}
}

func (em *ExipreMap) SetExpiredIn(k string, val interface{}, expiredIn time.Time) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.store[k] = Value{
		Val:       val,
		expiredIn: expiredIn,
	}
}

func (em *ExipreMap) Size() int {
	em.Gc()
	em.mu.RLock()
	defer em.mu.RUnlock()
	return len(em.store)
}

func (em *ExipreMap) ToMap() map[string]interface{} {
	em.Gc()
	em.mu.RLock()
	defer em.mu.RUnlock()
	res := make(map[string]interface{})
	for k, v := range em.store {
		res[k] = v.Val
	}
	return res
}

func (em *ExipreMap) Has(k string) bool {
	return !em.isExpired(k)
}

func (em *ExipreMap) Update(k string, newVal interface{}) error {
	if !em.Has(k) {
		return NO_KEY_TO_UPDATE
	}
	em.mu.Lock()
	defer em.mu.Unlock()
	val := em.store[k]
	val.Val = newVal
	em.store[k] = val
	return nil
}

func (em *ExipreMap) CleanUp() {
	if em.t != nil {
		em.t.Stop()
	}
}
