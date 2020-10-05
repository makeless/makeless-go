package makeless_go_config

type Config interface {
	Load(path string) error
	GetConfiguration() Configuration
}
