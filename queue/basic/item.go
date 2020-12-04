package basic

import "sync"

type Item struct {
	Data  interface{}
	Async bool
	*sync.RWMutex
}

func (item *Item) GetData() interface{} {
	item.RLock()
	defer item.RUnlock()

	return item.Data
}

func (item *Item) GetAsync() bool {
	item.RLock()
	defer item.RUnlock()

	return item.Async
}
