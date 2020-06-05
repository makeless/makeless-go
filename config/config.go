package go_saas_config

type Config interface {
	Load(path string) error
	GetConfiguration() Configuration
}
