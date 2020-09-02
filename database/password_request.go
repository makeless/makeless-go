package go_saas_database

import "github.com/go-saas/go-saas/model"

type PasswordRequest interface {
	CreatePasswordRequest(passwordRequest *go_saas_model.PasswordRequest) error
}
