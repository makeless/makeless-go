package go_saas_config

type Team interface {
	GetTokens() bool
	GetRoles() map[string]interface{}
	HasRole(role string) bool
}
