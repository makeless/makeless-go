package basic

import "sync"

type Node struct {
	Data interface{}

	next *Node
	*sync.RWMutex
}

func (node *Node) getNext() *Node {
	node.RLock()
	defer node.RUnlock()

	return node.next
}

func (node *Node) setNext(next *Node) {
	node.Lock()
	defer node.Unlock()

	node.next = next
}

func (node *Node) GetData() interface{} {
	node.RLock()
	defer node.RUnlock()

	return node.Data
}
