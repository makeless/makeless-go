package go_saas_event_basic

import (
	"github.com/gin-contrib/sse"
	"sync"
)

type Channel struct {
	Event sse.Event
	Error error
	*sync.RWMutex
}

func (channel *Channel) GetEvent() sse.Event {
	channel.RLock()
	defer channel.RUnlock()

	return channel.Event
}

func (channel *Channel) GetError() error {
	channel.RLock()
	defer channel.RUnlock()

	return channel.Error
}
