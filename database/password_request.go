package go_saas_database

import "github.com/go-saas/go-saas/model"

type PasswordRequest interface {
	CreatePasswordRequest(passwordRequest *go_saas_model.PasswordRequest) error
	GetPasswordRequest(passwordRequest *go_saas_model.PasswordRequest) (*go_saas_model.PasswordRequest, error)
	UpdatePasswordRequest(passwordRequest *go_saas_model.PasswordRequest) (*go_saas_model.PasswordRequest, error)
}
