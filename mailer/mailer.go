package go_saas_mailer

import "context"

type Mailer interface {
	Init() error
	Send(ctx context.Context, mail Mail) error
}
