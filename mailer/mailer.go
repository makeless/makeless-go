package makeless_go_mailer

import (
	"context"
	"github.com/makeless/makeless-go/queue"
)

type Mailer interface {
	Init() error
	GetHandlers() map[string]func(data map[string]interface{}, locale string) (Mail, error)
	GetHandler(name string) (func(data map[string]interface{}, locale string) (Mail, error), error)
	SetHandler(name string, handler func(data map[string]interface{}, locale string) (Mail, error))
	GetMail(name string, data map[string]interface{}, locale string) (Mail, error)
	GetQueue() makeless_go_queue.Queue
	Send(ctx context.Context, mail Mail) error
	SendQueue(mail Mail) error
}
