package makeless_go_tls

type Tls interface {
	GetCertPath() string
	GetKeyPath() string
}
