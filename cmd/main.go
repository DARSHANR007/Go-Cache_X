package main

import (
	"fmt"
	"go_cache/internals/cache"
	"time"
)

func main() {

	fmt.Println("hello")

	c := cache.NewCache()

	c.Set("Darshan", "7", 10*time.Second)

	value, found := c.GetItem("7")

	if found {
		fmt.Println("Found value:", *value)
	} else {
		fmt.Println("Value not found or expired.")
	}

	time.Sleep(11 * time.Second)

	value2, found2 := c.GetItem("7")

	if found2 {
		fmt.Println("Found value after expiration:", *value2)
	} else {
		fmt.Println("Value not found or expired after waiting.")
	}

}
