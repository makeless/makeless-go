package go_saas_mailer

type Mailer interface {
	Init() error
	Send(mail Mail) error
}
