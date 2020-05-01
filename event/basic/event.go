package saas_event_basic

import (
	"github.com/gin-contrib/sse"
	"github.com/loeffel-io/go-saas/event"
	"sync"
	"time"
)

type Event struct {
	Hub saas_event.Hub
	*sync.RWMutex
}

func (event *Event) NewClientId() uint {
	return uint(time.Now().Unix())
}

func (event *Event) GetHub() saas_event.Hub {
	event.RLock()
	defer event.RUnlock()

	return event.Hub
}

func (event *Event) Subscribe(userId uint, clientId uint) {
	event.GetHub().NewClient(userId, clientId)
}

func (event *Event) Unsubscribe(userId uint, clientId uint) {
	event.GetHub().DeleteClient(userId, clientId)
}

func (event *Event) Emit(userId uint, data sse.Event) {
	for _, channel := range event.GetHub().GetUser(userId) {
		channel <- data
	}
}

func (event *Event) Listen(userId uint, clientId uint) chan sse.Event {
	return event.GetHub().GetClient(userId, clientId)
}
