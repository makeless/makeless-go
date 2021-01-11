package basic

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPush(t *testing.T) {
	var queue = &Queue{
		Context: context.Background(),
		RWMutex: new(sync.RWMutex),
	}

	var firstNode = &Node{
		Data:    0,
		next:    nil,
		RWMutex: new(sync.RWMutex),
	}

	var secondNode = &Node{
		Data:    0,
		next:    firstNode,
		RWMutex: new(sync.RWMutex),
	}

	tests := []struct {
		Queue        *Queue
		Node         *Node
		Err          error
		ExpectedHead *Node
		ExpectedTail *Node
	}{
		{
			Queue:        queue,
			Node:         firstNode,
			Err:          nil,
			ExpectedHead: firstNode,
			ExpectedTail: firstNode,
		},
		{
			Queue:        queue,
			Node:         secondNode,
			Err:          nil,
			ExpectedHead: firstNode,
			ExpectedTail: secondNode,
		},
	}

	for i, test := range tests {
		err := test.Queue.Add(test.Node)
		assert.Equalf(t, test.Err, err, "%d: error not equal", i)

		head := test.Queue.getHead()
		tail := test.Queue.getTail()

		assert.Equal(t, test.ExpectedHead, head)
		assert.Equal(t, test.ExpectedTail, tail)
	}
}
