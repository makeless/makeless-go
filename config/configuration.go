package go_saas_config

type Configuration interface {
	GetName() string
	GetLogo() string
	GetLocale() string
	GetHost() string
	GetMail() Mail
	GetTokens() bool
	GetTeams() Team
}
