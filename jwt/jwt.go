package go_saas_jwt

type Jwt interface {
	GetId() string
	GetKey() string
}
