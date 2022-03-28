package makeless_go_makeless_basic

import (
	"github.com/makeless/makeless-go/config"
	"github.com/makeless/makeless-go/database/database"
	"github.com/makeless/makeless-go/logger"
	"github.com/makeless/makeless-go/mailer"
	"gorm.io/gorm"
	"sync"
)

type Makeless struct {
	Config   makeless_go_config.Config
	Logger   makeless_go_logger.Logger
	Mailer   makeless_go_mailer.Mailer
	Database makeless_go_database.Database
	*sync.RWMutex
}

func (makeless *Makeless) GetConfig() makeless_go_config.Config {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Config
}

func (makeless *Makeless) GetLogger() makeless_go_logger.Logger {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Logger
}

func (makeless *Makeless) GetMailer() makeless_go_mailer.Mailer {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Mailer
}

func (makeless *Makeless) GetDatabase() makeless_go_database.Database {
	makeless.RLock()
	defer makeless.RUnlock()

	return makeless.Database
}

func (makeless *Makeless) Init(dialector gorm.Dialector, path string) error {
	if err := makeless.GetConfig().Load(path); err != nil {
		return err
	}

	if err := makeless.GetMailer().GetQueue().Init(); err != nil {
		return err
	}

	if err := makeless.GetMailer().Init(); err != nil {
		return err
	}

	if err := makeless.GetDatabase().Connect(dialector); err != nil {
		return err
	}

	if err := makeless.GetDatabase().Migrate(); err != nil {
		return err
	}

	return nil
}
