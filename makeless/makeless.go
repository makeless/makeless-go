package makeless

import (
	"github.com/makeless/makeless-go/v2/config"
	"github.com/makeless/makeless-go/v2/database/database"
	"github.com/makeless/makeless-go/v2/logger"
	"github.com/makeless/makeless-go/v2/mailer"
	"gorm.io/gorm"
)

type Makeless interface {
	GetConfig() makeless_go_config.Config
	GetLogger() makeless_go_logger.Logger
	GetMailer() makeless_go_mailer.Mailer
	GetDatabase() makeless_go_database.Database
	Init(dialector gorm.Dialector, path string) error
}
