package basic

import (
	"context"
	"sync"
)

type Queue struct {
	Context context.Context

	*sync.Map
	*sync.RWMutex
}

func (queue *Queue) GetContext() context.Context {
	queue.RLock()
	defer queue.RUnlock()

	return queue.Context
}

func (queue *Queue) Push(channel string, item *Item) error {
	list, exists := queue.Load(channel)

	if list == nil || !exists {
		list = make([]*Item, 0)
	}

	list = append(list.([]*Item), item)
	queue.Store(channel, list)
	return nil
}

func (queue *Queue) Pop(channel string) (*Item, error) {
	list, exists := queue.Load(channel)

	if list == nil || !exists {
		return nil, nil
	}

	return list.([]*Item)[len(list.([]*Item))-1], nil
}
