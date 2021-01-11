package makeless_go_queue_basic

import (
	"context"
	"github.com/makeless/makeless-go/queue"
	"sync"
)

type Queue struct {
	Context context.Context

	head makeless_go_queue.Node
	tail makeless_go_queue.Node
	*sync.RWMutex
}

func (queue *Queue) getHead() makeless_go_queue.Node {
	queue.RLock()
	defer queue.RUnlock()

	return queue.head
}

func (queue *Queue) setHead(head makeless_go_queue.Node) {
	queue.Lock()
	defer queue.Unlock()

	queue.head = head
}

func (queue *Queue) getTail() makeless_go_queue.Node {
	queue.RLock()
	defer queue.RUnlock()

	return queue.tail
}

func (queue *Queue) setTail(tail makeless_go_queue.Node) {
	queue.Lock()
	defer queue.Unlock()

	queue.tail = tail
}

func (queue *Queue) GetContext() context.Context {
	queue.RLock()
	defer queue.RUnlock()

	return queue.Context
}

func (queue *Queue) Add(node makeless_go_queue.Node) error {
	if queue.getTail() != nil {
		queue.getTail().SetNext(node)
	}

	queue.setTail(node)

	if queue.getHead() == nil {
		queue.setHead(node)
	}

	return nil
}

func (queue *Queue) Remove() (makeless_go_queue.Node, error) {
	var head = queue.getHead()

	if head == nil {
		return nil, nil
	}

	var next = head.GetNext()

	queue.setHead(next)

	if queue.getHead() == nil {
		queue.setTail(nil)
	}

	return head, nil
}

func (queue *Queue) Empty() bool {
	return queue.getHead() == nil
}
