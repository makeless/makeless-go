package makeless_go_queue

import "context"

type Queue interface {
	GetContext() context.Context
	Pop(channel string, item Item) error
	Push(channel string) (Item, error)
}
