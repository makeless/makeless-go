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
	event.Lock()
	defer event.Unlock()

	event.Hub.NewClient(userId, clientId)
}

func (event *Event) Unsubscribe(userId uint, clientId uint) {
	event.Lock()
	defer event.Unlock()

	event.Hub.DeleteClient(userId, clientId)
}

func (event *Event) Trigger(userId uint, channel string, id string, data interface{}) {
	event.Lock()
	defer event.Unlock()

	for _, client := range event.Hub.GetUser(userId) {
		client <- sse.Event{
			Event: channel,
			Retry: 3,
			Data: Data{
				Id:   id,
				Data: data,
			},
		}
	}
}

func (event *Event) Broadcast(channel string, id string, data interface{}) {
	for userId := range event.GetHub().GetList() {
		event.Trigger(userId, channel, id, data)
	}
}

func (event *Event) Listen(userId uint, clientId uint) chan sse.Event {
	event.Lock()
	defer event.Unlock()

	return event.Hub.GetClient(userId, clientId)
}
