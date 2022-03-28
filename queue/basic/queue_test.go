package makeless_go_queue_basic

import (
	"context"
	"github.com/makeless/makeless-go/v2/queue"
	"reflect"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	var queue = &Queue{
		Context: context.Background(),
		RWMutex: new(sync.RWMutex),
	}

	var firstNode = &Node{
		Data:    []byte("0"),
		next:    nil,
		RWMutex: new(sync.RWMutex),
	}

	var secondNode = &Node{
		Data:    []byte("1"),
		next:    firstNode,
		RWMutex: new(sync.RWMutex),
	}

	var thirdNode = &Node{
		Data:    []byte("2"),
		next:    secondNode,
		RWMutex: new(sync.RWMutex),
	}

	tests := []struct {
		queue        *Queue
		node         makeless_go_queue.Node
		err          error
		expectedHead makeless_go_queue.Node
		expectedTail makeless_go_queue.Node
	}{
		{
			queue:        queue,
			node:         firstNode,
			err:          nil,
			expectedHead: firstNode,
			expectedTail: firstNode,
		},
		{
			queue:        queue,
			node:         secondNode,
			err:          nil,
			expectedHead: firstNode,
			expectedTail: secondNode,
		},
		{
			queue:        queue,
			node:         thirdNode,
			err:          nil,
			expectedHead: firstNode,
			expectedTail: thirdNode,
		},
	}

	for i, test := range tests {
		err := test.queue.Add(test.node)

		if err != nil {
			t.Fatalf("%d: error not equal", i)
		}

		head := test.queue.getHead()

		if !reflect.DeepEqual(head, test.expectedHead) {
			t.Fatalf("%d: head not equal", i)
		}

		tail := test.queue.getTail()

		if !reflect.DeepEqual(tail, test.expectedTail) {
			t.Fatalf("%d: tail not equal", i)
		}
	}
}
