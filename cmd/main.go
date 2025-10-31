package main

import (
	"fmt"
	LRU "go_cache/evictionPolicies/LRU"
)

func main() {
	cache := LRU.New(3)

	cache.Put("A", 1)
	cache.Put("B", 2)
	cache.Put("C", 3)
	cache.Display()

	cache.Get("A") // Access A → now A is most recent
	cache.Display()

	cache.Put("D", 4) // Should evict least recently used (B)
	cache.Display()

	cache.Put("E", 5) // Should evict least recently used (C)
	cache.Display()

	if val, ok := cache.Get("A"); ok {
		fmt.Println("✅ Got A:", val)
	} else {
		fmt.Println("❌ A not found")
	}
}
