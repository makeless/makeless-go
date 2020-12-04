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
		Map:     new(sync.Map),
		RWMutex: new(sync.RWMutex),
	}

	tests := []struct {
		Queue    *Queue
		Channel  string
		Item     *Item
		Err      error
		Expected []*Item
	}{
		{
			Channel: "numbers",
			Queue:   queue,
			Item: &Item{
				Data:  0,
				Async: false,
			},
			Err: nil,
			Expected: []*Item{
				{
					Data:  0,
					Async: false,
				},
			},
		},
		{
			Channel: "numbers",
			Queue:   queue,
			Item: &Item{
				Data:  1,
				Async: false,
			},
			Err: nil,
			Expected: []*Item{
				{
					Data:  0,
					Async: false,
				},
				{
					Data:  1,
					Async: false,
				},
			},
		},
		{
			Channel: "strings",
			Queue:   queue,
			Item: &Item{
				Data:  "0",
				Async: false,
			},
			Err: nil,
			Expected: []*Item{
				{
					Data:  "0",
					Async: false,
				},
			},
		},
	}

	for i, test := range tests {
		err := test.Queue.Push(test.Channel, test.Item)
		assert.Equalf(t, test.Err, err, "%d: error not equal", i)

		load, _ := test.Queue.Load(test.Channel)
		assert.Equal(t, test.Expected, load)
	}
}
