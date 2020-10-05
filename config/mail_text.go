package makeless_go_config

type MailText interface {
	GetGreeting() string
	GetSignature() string
	GetCopyright() string
}
