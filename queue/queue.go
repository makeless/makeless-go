package makeless_go_queue

import "context"

type Queue interface {
	Init() error
	GetContext() context.Context
	Add(node Node) error
	Remove() (Node, error)
	Empty() (bool, error)
}
