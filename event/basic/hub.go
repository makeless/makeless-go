package makeless_go_event_basic

import (
	"github.com/gin-contrib/sse"
	"sync"
)

type Hub struct {
	List *sync.Map
	*sync.RWMutex
}

func (hub *Hub) GetList() *sync.Map {
	hub.RLock()
	defer hub.RUnlock()

	return hub.List
}

func (hub *Hub) GetUser(userId uint) *sync.Map {
	if user, exists := hub.GetList().Load(userId); exists {
		return user.(*sync.Map)
	}

	return nil
}

func (hub *Hub) GetClient(userId uint, clientId uint) chan sse.Event {
	var user = hub.GetUser(userId)
	var client interface{}

	if user == nil {
		return nil
	}

	if client, _ = user.Load(clientId); client == nil {
		return nil
	}

	return client.(chan sse.Event)
}

func (hub *Hub) DeleteClient(userId uint, clientId uint) {
	var user = hub.GetUser(userId)
	var client = hub.GetClient(userId, clientId)

	if user == nil || client == nil {
		return
	}

	close(client)
	user.Delete(clientId)
}

func (hub *Hub) NewClient(userId uint, clientId uint) {
	var user, _ = hub.GetList().LoadOrStore(userId, new(sync.Map))
	user.(*sync.Map).Store(clientId, make(chan sse.Event))
}
