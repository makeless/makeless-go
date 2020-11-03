package makeless_go_database

import (
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
	"gorm.io/gorm"
)

type Profile interface {
	UpdateProfile(connection *gorm.DB, user *makeless_go_model.User, profile *_struct.Profile) (*makeless_go_model.User, error)
	UpdateProfileTeam(connection *gorm.DB, team *makeless_go_model.Team) (*makeless_go_model.Team, error)
}
