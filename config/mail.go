package makeless_go_config

type Mail interface {
	GetName() string
	GetLogo() string
	GetFrom() string
	GetLink() string
	GetButtonColor() string
	GetButtonTextColor() string
	GetTexts(locale string) MailText
}
