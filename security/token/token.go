package makeless_go_security_token

type SecurityToken interface {
	Generate(length int) (string, error)
}
