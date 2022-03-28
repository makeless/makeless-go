package makeless_go_mailer

import (
	"context"
	"github.com/makeless/makeless-go/v2/mail"
	"github.com/makeless/makeless-go/v2/queue"
)

type Mailer interface {
	Init() error
	GetQueue() makeless_go_queue.Queue
	Send(ctx context.Context, mail makeless_go_mail.Mail) error
	SendQueue(mail makeless_go_mail.Mail) error
}
