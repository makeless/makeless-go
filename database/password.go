package go_saas_database

import "github.com/go-saas/go-saas/model"

type Password interface {
	UpdatePassword(user *go_saas_model.User, newPassword string) (*go_saas_model.User, error)
}
