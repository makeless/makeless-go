package go_saas_event_basic

import (
	"github.com/gin-contrib/sse"
	"github.com/go-saas/go-saas/event"
	"sync"
	"time"
)

type Event struct {
	Hub go_saas_event.Hub
	*sync.RWMutex
}

func (event *Event) GetHub() go_saas_event.Hub {
	event.RLock()
	defer event.RUnlock()

	return event.Hub
}

func (event *Event) NewClientId() uint {
	return uint(time.Now().Unix())
}

func (event *Event) Subscribe(userId uint, clientId uint) {
	event.GetHub().NewClient(userId, clientId)
}

func (event *Event) Unsubscribe(userId uint, clientId uint) {
	event.GetHub().DeleteClient(userId, clientId)
}

func (event *Event) Trigger(userId uint, channel string, id string, data interface{}) {
	event.GetHub().GetUser(userId).Range(func(clientId, client interface{}) bool {
		client.(chan sse.Event) <- sse.Event{
			Event: channel,
			Retry: 3,
			Data: Data{
				Id:      id,
				Data:    data,
				RWMutex: new(sync.RWMutex),
			},
		}

		return true
	})
}

func (event *Event) Broadcast(channel string, id string, data interface{}) {
}

func (event *Event) Listen(userId uint, clientId uint) chan sse.Event {
	return event.GetHub().GetClient(userId, clientId)
}
