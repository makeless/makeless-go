package go_saas_config_basic

import (
	"encoding/json"
	"github.com/go-saas/go-saas/config"
	"io/ioutil"
	"os"
	"sync"
)

type Config struct {
	Configuration go_saas_config.Configuration
	*sync.RWMutex
}

func (config *Config) Load(path string) error {
	config.Lock()
	defer config.Unlock()

	var err error
	var bytes []byte
	var file *os.File

	var team = &Team{
		RWMutex: new(sync.RWMutex),
	}

	var configuration = &Configuration{
		RWMutex: new(sync.RWMutex),
		Teams:   team,
	}

	if file, err = os.Open(path); err != nil {
		return err
	}

	if bytes, err = ioutil.ReadAll(file); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &configuration); err != nil {
		return err
	}

	config.Configuration = configuration
	return nil
}

func (config *Config) GetConfiguration() go_saas_config.Configuration {
	config.RLock()
	defer config.RUnlock()

	return config.Configuration
}
