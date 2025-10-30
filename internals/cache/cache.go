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

	return c

}

func (c *Cache) Set(value interface{}, key string, ttl time.Duration) {

	c.mu.Lock()
	defer c.mu.Unlock()

	var expTime int64

	if ttl > 0 {
		expTime = time.Now().Add(ttl).Unix()
	}

	c.data[key] = Item{CacheValue: value, Expiration: expTime}

}

func (c *Cache) GetItem(key string) (*interface{}, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.data[key]

	if !found {
		fmt.Println("Error in retrieving")
		fmt.Errorf("Item not found")
		return nil, false

	}

	if item.Expiration > 0 && time.Now().Unix() > item.Expiration {

		fmt.Errorf(" Item has been expired")

		return nil, false
	}

	return &item.CacheValue, true

}

func (c *Cache) DeleteCache(key string) {

	c.mu.Lock()
	defer c.mu.Lock()
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
			deletedKeys = append(deletedKeys, key)
			delete(c.data, key)
		}
	}

	return deletedKeys
}

func (c *Cache) autoExpiration() {

	heartbeat := time.NewTicker(c.cleanupPeriod)
	defer heartbeat.Stop()

}

func (c *Cache) stopCleanup() {
	close(c.stopCleanUp)
}
