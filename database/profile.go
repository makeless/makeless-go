package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
)

type Profile interface {
	UpdateProfile(user *go_saas_model.User, profile *_struct.Profile) (*go_saas_model.User, error)
	UpdateProfileTeam(team *go_saas_model.Team) (*go_saas_model.Team, error)
}
