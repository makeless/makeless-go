package go_saas_event

import "github.com/gin-contrib/sse"

type Channel interface {
	GetEvent() sse.Event
	GetError() error
}
