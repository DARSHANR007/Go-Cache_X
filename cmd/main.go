package main

import (
	"fmt"
	"go_cache/internals/cache"
	"time"
)

func main() {
	c := cache.NewCache(3 * time.Second)

	c.Set("user1", "Darsh", 5*time.Second)
	c.Set("user2", "Alice", 2*time.Second)
	c.Set("user3", "Bob", 7*time.Second)
	c.Set("user4", "Charlie", 10*time.Second)
	c.Set("user5", "Eve", 1*time.Second)
	c.Set("user6", "Mallory", 8*time.Second)
	c.Set("user7", "Trudy", 3*time.Second)
	c.Set("user8", "Victor", 4*time.Second)
	c.Set("user9", "Peggy", 12*time.Second)
	c.Set("user10", "Sybil", 6*time.Second)

	fmt.Println("✅ Added 10 users to cache")

	time.Sleep(15 * time.Second)

	fmt.Println("\n🔍 Checking who is still alive in cache:")

	for i := 1; i <= 10; i++ {
		key := fmt.Sprintf("user%d", i)
		if val, ok := c.GetItem(key); ok {
			fmt.Printf("✅ %s found: %v\n", key, val)
		} else {
			fmt.Printf(" %s expired or missing\n", key)
		}
	}

	c.FinishCleanup()
}
