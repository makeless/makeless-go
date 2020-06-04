package go_saas_basic_security

import "github.com/go-saas/go-saas/database"

func (security *Security) getDatabase() *saas_database.Database {
	security.RLock()
	defer security.RUnlock()

	return security.Database
}
