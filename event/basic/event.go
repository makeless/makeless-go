package makeless_go_event_basic

import (
	"github.com/google/uuid"
	"github.com/makeless/makeless-go/event"
	"sync"
)

type Event struct {
	Name  string
	Hub   makeless_go_event.Hub
	Error chan error
	*sync.RWMutex
}

func (event *Event) Init() error {
	return nil
}

func (event *Event) GetName() string {
	event.RLock()
	defer event.RUnlock()

	return event.Name
}

func (event *Event) GetHub() makeless_go_event.Hub {
	event.RLock()
	defer event.RUnlock()

	return event.Hub
}

func (event *Event) GetError() chan error {
	event.RLock()
	defer event.RUnlock()

	return event.Error
}

func (event *Event) NewClientId() string {
	return uuid.NewString()
}

func (event *Event) Subscribe(userId uint, clientId string) {
	event.GetHub().NewClient(userId, clientId)
}

func (event *Event) Unsubscribe(userId uint, clientId string) {
	event.GetHub().DeleteClient(userId, clientId)
}

func (event *Event) Trigger(userId uint, channel string, id string, data interface{}) error {
	var user = event.GetHub().GetUser(userId)

	if user == nil {
		return nil
	}

	user.Range(func(clientId, client interface{}) bool {
		client.(chan makeless_go_event.EventData) <- &EventData{
			Channel: channel,
			Id:      id,
			Data:    data,
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

func (event *Event) Listen(userId uint, clientId string) chan makeless_go_event.EventData {
	return event.GetHub().GetClient(userId, clientId)
}

func (event *Event) ListenError() chan error {
	return event.GetError()
}
