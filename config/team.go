package makeless_go_config

type Team interface {
	GetTokens() bool
	GetRoles() map[string]interface{}
	HasRole(role string) bool
}
