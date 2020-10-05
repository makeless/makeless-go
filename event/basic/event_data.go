package makeless_go_event_basic

import "sync"

type EventData struct {
	Id   string      `json:"id"`
	Data interface{} `json:"data"`

	*sync.RWMutex
}

func (eventData *EventData) GetId() string {
	eventData.RLock()
	defer eventData.RUnlock()

	return eventData.Id
}

func (eventData *EventData) GetData() interface{} {
	eventData.RLock()
	defer eventData.RUnlock()

	return eventData.Data
}
