package makeless

import (
	"github.com/makeless/makeless-go/config"
	"github.com/makeless/makeless-go/database/database"
	"github.com/makeless/makeless-go/logger"
	"github.com/makeless/makeless-go/mailer"
	"gorm.io/gorm"
)

type Makeless interface {
	GetConfig() makeless_go_config.Config
	GetLogger() makeless_go_logger.Logger
	GetMailer() makeless_go_mailer.Mailer
	GetDatabase() makeless_go_database.Database
	Init(dialector gorm.Dialector, path string) error
}
