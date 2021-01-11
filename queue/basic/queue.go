package basic

import (
	"context"
	"sync"
)

type Queue struct {
	Context context.Context

	head *Node
	tail *Node
	*sync.RWMutex
}

func (queue *Queue) getHead() *Node {
	queue.RLock()
	defer queue.RUnlock()

	return queue.head
}

func (queue *Queue) setHead(head *Node) {
	queue.Lock()
	defer queue.Unlock()

	queue.head = head
}

func (queue *Queue) getTail() *Node {
	queue.RLock()
	defer queue.RUnlock()

	return queue.tail
}

func (queue *Queue) setTail(tail *Node) {
	queue.Lock()
	defer queue.Unlock()

	queue.tail = tail
}

func (queue *Queue) GetContext() context.Context {
	queue.RLock()
	defer queue.RUnlock()

	return queue.Context
}

func (queue *Queue) Add(node *Node) error {
	if queue.getTail() != nil {
		queue.getTail().setNext(node)
	}

	queue.setTail(node)

	if queue.getHead() == nil {
		queue.setHead(node)
	}

	return nil
}

func (queue *Queue) Remove() (*Node, error) {
	return nil, nil
}

func (queue *Queue) Empty() bool {
	return queue.getHead() == nil
}
