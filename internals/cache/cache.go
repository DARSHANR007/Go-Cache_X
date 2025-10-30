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
	stopCleanUp   chan struct{}
}

// Constructor
func NewCache(cleanupPeriod time.Duration) *Cache {
	c := &Cache{
		data:          make(map[string]Item),
		cleanupPeriod: cleanupPeriod,
		stopCleanUp:   make(chan struct{}),
	}
	go c.autoExpiration() // start cleanup in background
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
		if item.Expiration > 0 && item.Expiration <= currTime {
			fmt.Printf("🧹 Deleted expired key: %s\n", key)
			deletedKeys = append(deletedKeys, key)
			delete(c.data, key)
		}
	}

	return deletedKeys
}

func (c *Cache) autoExpiration() {
	ticker := time.NewTicker(c.cleanupPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.autoCleanUp()
		case <-c.stopCleanUp:
			return
		}
	}
}

func (c *Cache) FinishCleanup() {
	close(c.stopCleanUp)
}
