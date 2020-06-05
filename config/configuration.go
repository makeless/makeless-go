package go_saas_config

type Configuration interface {
	GetName() string
	GetTokens() bool
	GetTeams() Team
}
