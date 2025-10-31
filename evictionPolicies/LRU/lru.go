package lru

import "fmt"

type Cache struct {
	capacity int
	items    map[string]*Node
	list     *DoublyLinkedList
}

func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		items:    make(map[string]*Node),
		list:     &DoublyLinkedList{},
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if node, found := c.items[key]; found {
		c.list.MoveToFront(node)
		fmt.Println("Successfull cache hit")
		return node.value, true
	}

	return nil, false
}

func (c *Cache) Set(key string, value interface{}) bool {

	if node, found := c.items[key]; found { //update, set new
		node.value = value
		c.list.MoveToFront(node)
		fmt.Println("key already present,updated")
		return true
	}

	if len(c.items) >= c.capacity {

		removed := c.list.RemoveFromTail()
		delete(c.items, removed.key)

	}

	newNode := &Node{key: key, value: value}
	c.list.AddtoFront(newNode)
	c.items[key] = newNode
	fmt.Println("added new node")
	return true

}

func (c *Cache) Display() {
	fmt.Print("Cache state (MRU → LRU): ")
	for node := c.list.head; node != nil; node = node.next {
		fmt.Printf("[%s:%v] ", node.key, node.value)
	}
	fmt.Println()
}
