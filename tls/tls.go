package go_saas_tls

type Tls interface {
	GetCertPath() string
	GetKeyPath() string
}
