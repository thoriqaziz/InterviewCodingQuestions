// Goroutine-safe Cache with TTL

// Problem without caching -> High DB Load, Latency Increase, Pool Scalability, High Infra Cost
// Prevent Race Condition
// sync.Mutex
// Sync.RWMutex -> Multiple readers at the same time, only 1 writer and no readers during write

package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]Item
}

type Item struct {
	value      string
	expiration int64
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]Item),
	}
}

func (c *Cache) Set(key, val string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = Item{
		value:      val,
		expiration: time.Now().Add(ttl).Unix(),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.data[key]
	if !ok {
		return "", false
	}
	if time.Now().Unix() > item.expiration {
		return "", false
	}
	return item.value, true
}

func main() {
	cache := NewCache()
	cache.Set("test1", "Thoriq Aziz", 3*time.Second)

	val, found := cache.Get("test1")
	fmt.Println(val, found)

	time.Sleep(4 * time.Second)
	val1, found1 := cache.Get("test1")
	fmt.Println(val1, found1)
}
