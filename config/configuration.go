package makeless_go_config

type Configuration interface {
	GetName() string
	GetLogo() string
	GetLocale() string
	GetHost() string
	GetEmailVerification() bool
	GetTokens() bool
	GetTeams() Team
	GetMail() Mail
}
