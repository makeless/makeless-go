package makeless_go_config_basic

import (
	"encoding/json"
	"github.com/makeless/makeless-go/config"
	"io/ioutil"
	"os"
	"sync"
)

type Config struct {
	Configuration makeless_go_config.Configuration
	*sync.RWMutex
}

func (config *Config) Load(path string) error {
	config.Lock()
	defer config.Unlock()

	var err error
	var bytes []byte
	var file *os.File

	var mail = &Mail{
		Texts:   make(map[string]*MailText),
		RWMutex: new(sync.RWMutex),
	}

	var team = &Team{
		RWMutex: new(sync.RWMutex),
	}

	var configuration = &Configuration{
		Mail:    mail,
		Teams:   team,
		RWMutex: new(sync.RWMutex),
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

	for locale := range mail.Texts {
		mail.Texts[locale].RWMutex = new(sync.RWMutex)
	}

	config.Configuration = configuration
	return nil
}

func (config *Config) GetConfiguration() makeless_go_config.Configuration {
	config.RLock()
	defer config.RUnlock()

	return config.Configuration
}
