package cache

import (
	"sync"
	"sync/atomic"
	"time"
)

type (
	cacheNode[V any] struct {
		data       V
		createTime time.Time
		accessTime time.Time
		hits       atomic.Int32
	}

	Cache[K, V any] struct {
		m   sync.Map
		cap int
	}
)

func New[K, V any](cap int) *Cache[K, V] {
	return &Cache[K, V]{
		cap: cap,
	}
}

func (p *Cache[K, V]) Get(key K) (v V, ok bool) {
	var value any
	if value, ok = p.m.Load(key); ok {
		node := value.(*cacheNode[V])

		v = node.data

		node.accessTime = time.Now()
		node.hits.Add(1)
	}

	return
}

func (p *Cache[K, V]) Set(key K, value V) {
	now := time.Now()
	p.m.Store(key, &cacheNode[V]{
		data:       value,
		createTime: now,
		accessTime: now,
		hits:       atomic.Int32{},
	})
}

func (p *Cache[K, V]) Clear() {
	p.m.Clear()
}
