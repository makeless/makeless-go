package saas_event_basic

import (
	"github.com/gin-contrib/sse"
	"github.com/go-saas/go-saas/event"
	"sync"
)

type Hub struct {
	List map[uint]map[uint]chan sse.Event
	*sync.RWMutex
}

func (hub *Hub) Init() saas_event.Hub {
	hub.List = make(map[uint]map[uint]chan sse.Event)
	hub.RWMutex = new(sync.RWMutex)

	return hub
}

func (hub *Hub) GetList() map[uint]map[uint]chan sse.Event {
	hub.RLock()
	defer hub.RUnlock()

	return hub.List
}

func (hub *Hub) GetUser(userId uint) map[uint]chan sse.Event {
	hub.RLock()
	defer hub.RUnlock()

	if user, exists := hub.List[userId]; exists {
		return user
	}

	return nil
}

func (hub *Hub) GetClient(userId uint, clientId uint) chan sse.Event {
	if hub.GetUser(userId) == nil {
		return nil
	}

	hub.RLock()
	defer hub.RUnlock()

	return hub.List[userId][clientId]
}

func (hub *Hub) DeleteClient(userId uint, clientId uint) {
	if hub.GetClient(userId, clientId) == nil {
		return
	}

	hub.Lock()
	defer hub.Unlock()

	close(hub.List[userId][clientId])
	delete(hub.List[userId], clientId)
}

func (hub *Hub) NewClient(userId uint, clientId uint) {
	if hub.GetUser(userId) == nil {
		hub.Lock()
		hub.List[userId] = make(map[uint]chan sse.Event)
		hub.Unlock()
	}

	hub.Lock()
	defer hub.Unlock()
	hub.List[userId][clientId] = make(chan sse.Event)
}
