package go_saas_config

type MailText interface {
	GetGreeting() string
	GetSignature() string
	GetCopyright() string
}
