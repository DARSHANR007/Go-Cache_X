package cache

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	CacheValue interface{}
	Expiration int64
}

type Cache struct {
	data          map[string]Item
	mu            sync.RWMutex
	cleanupPeriod time.Duration
	stopCleanUp   chan struct{} //empty channel to notify
}

// Constructor of cache
func NewCache(cleanupPeriod time.Duration) *Cache {

	c := &Cache{
		data: make(map[string]Item), cleanupPeriod: cleanupPeriod, stopCleanUp: make(chan struct{}),
	}

	go c.autoExpiration()

	return c

}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {

	c.mu.Lock()
	defer c.mu.Unlock()

	var expTime int64

	if ttl > 0 {
		expTime = time.Now().Add(ttl).Unix()
	}

	c.data[key] = Item{CacheValue: value, Expiration: expTime}

}

func (c *Cache) GetItem(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.data[key]
	if !found {
		fmt.Println("❌ Item not found")
		return nil, false
	}

	if item.Expiration > 0 && time.Now().Unix() > item.Expiration {
		fmt.Println("⚠️ Item expired")
		return nil, false
	}

	return item.CacheValue, true
}

func (c *Cache) DeleteCache(key string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)

}

func (c *Cache) autoCleanUp() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	currTime := time.Now().Unix()
	var deletedKeys []string

	for key, item := range c.data {
		// check if item has an expiration and is expired
		if item.Expiration > 0 && item.Expiration <= currTime {
			fmt.Printf("🧹 Deleted expired key: %s\n", k)

			deletedKeys = append(deletedKeys, key)
			delete(c.data, key)
		}
	}

	return deletedKeys
}

func (c *Cache) autoExpiration() {

	heartbeat := time.NewTicker(c.cleanupPeriod)
	defer heartbeat.Stop()

	for {
		select {
		case <-heartbeat.C: // C in inbuilt inside ticker
			c.autoCleanUp()
		case <-c.stopCleanUp:
			return
		}
	}

}

func (c *Cache) stopCleanup() {
	close(c.stopCleanUp)
}
