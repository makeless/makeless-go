package go_saas_config

type Configuration interface {
	GetName() string
	GetLogo() string
	GetLocale() string
	GetHost() string
	GetMail() string
	GetTokens() bool
	GetTeams() Team
}
