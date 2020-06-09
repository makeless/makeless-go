package go_saas_database

import "github.com/go-saas/go-saas/model"

type User interface {
	GetUser(user *go_saas_model.User) (*go_saas_model.User, error)
	GetUserByField(user *go_saas_model.User, field string, value string) (*go_saas_model.User, error)
	CreateUser(user *go_saas_model.User) (*go_saas_model.User, error)
}
