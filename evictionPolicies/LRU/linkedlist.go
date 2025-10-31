package lru

type Node struct {
	key   string
	value interface{}
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) AddtoFront(currNode *Node) {
	currNode.prev = nil
	currNode.next = dll.head // add the most recently used to the left most ie head

	if dll.head != nil {
		dll.head.prev = currNode
	}
	dll.head = currNode

	if dll.tail == nil {
		dll.tail = currNode
	}

}

func (dll *DoublyLinkedList) RemoveFromList(currNode *Node) {

	if currNode.prev != nil {
		currNode.prev.next = currNode.next
	} else {
		dll.head = currNode.next
	}

	if currNode.next != nil {
		currNode.next.prev = currNode.prev

	} else {
		dll.tail = currNode.prev
	}
}

func (dll *DoublyLinkedList) MoveToFront(currNode *Node) {
	if dll.head == currNode {
		return
	}
	dll.RemoveFromList(currNode)
	dll.AddtoFront(currNode)
}

func (dll *DoublyLinkedList) RemoveFromTail() *Node {
	if dll.tail == nil {
		return nil
	}
	node := dll.tail
	dll.RemoveFromList(node)
	return node
}
