package go_saas_event_basic

import (
	"github.com/gin-contrib/sse"
	"github.com/go-saas/go-saas/event"
	"sync"
	"time"
)

type Event struct {
	Hub   go_saas_event.Hub
	Error chan error
	*sync.RWMutex
}

func (event *Event) Init() error {
	return nil
}

func (event *Event) GetHub() go_saas_event.Hub {
	event.RLock()
	defer event.RUnlock()

	return event.Hub
}

func (event *Event) GetError() chan error {
	event.RLock()
	defer event.RUnlock()

	return event.Error
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

func (event *Event) Trigger(userId uint, channel string, id string, data interface{}) error {
	event.GetHub().GetUser(userId).Range(func(clientId, client interface{}) bool {
		client.(chan sse.Event) <- sse.Event{
			Event: channel,
			Retry: 3,
			Data: &EventData{
				Id:      id,
				Data:    data,
				RWMutex: new(sync.RWMutex),
			},
		}

		return true
	})

	return nil
}

func (event *Event) TriggerError(err error) {
	event.GetError() <- err
}

func (event *Event) Broadcast(channel string, id string, data interface{}) error {
	var err error

	event.GetHub().GetList().Range(func(userId, value interface{}) bool {
		if err = event.Trigger(userId.(uint), channel, id, data); err != nil {
			return false
		}

		return true
	})

	return err
}

func (event *Event) Listen(userId uint, clientId uint) chan sse.Event {
	return event.GetHub().GetClient(userId, clientId)
}

func (event *Event) ListenError() chan error {
	return event.GetError()
}
