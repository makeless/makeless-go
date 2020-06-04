package go_saas_basic_security

import (
	"github.com/go-saas/go-saas/database/go_saas_basic_database"
)

func (security *Security) getDatabase() *go_saas_basic_database.saas_database {
	security.RLock()
	defer security.RUnlock()

	return security.Database
}
