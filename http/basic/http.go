package go_saas_http_basic

import (
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/authenticator"
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/event"
	"github.com/go-saas/go-saas/http"
	"github.com/go-saas/go-saas/logger"
	"github.com/go-saas/go-saas/security"
	"github.com/go-saas/go-saas/tls"
	"sync"
)

type Http struct {
	Router        *gin.Engine
	Handlers      map[string]func(http go_saas_http.Http) error
	Logger        go_saas_logger.Logger
	Event         go_saas_event.Event
	Authenticator go_saas_authenticator.Authenticator
	Security      go_saas_security.Security
	Database      go_saas_database.Database
	Tls           go_saas_tls.Tls
	Origins       []string
	Headers       []string
	Port          string
	Mode          string
	*sync.RWMutex
}

func (http *Http) GetRouter() *gin.Engine {
	http.RLock()
	defer http.RUnlock()

	return http.Router
}

func (http *Http) GetHandlers() map[string]func(http go_saas_http.Http) error {
	http.RLock()
	defer http.RUnlock()

	return http.Handlers
}

func (http *Http) GetLogger() go_saas_logger.Logger {
	http.RLock()
	defer http.RUnlock()

	return http.Logger
}

func (http *Http) GetEvent() go_saas_event.Event {
	http.RLock()
	defer http.RUnlock()

	return http.Event
}

func (http *Http) GetAuthenticator() go_saas_authenticator.Authenticator {
	http.RLock()
	defer http.RUnlock()

	return http.Authenticator
}

func (http *Http) GetSecurity() go_saas_security.Security {
	http.RLock()
	defer http.RUnlock()

	return http.Security
}

func (http *Http) GetDatabase() go_saas_database.Database {
	http.RLock()
	defer http.RUnlock()

	return http.Database
}

func (http *Http) GetTls() go_saas_tls.Tls {
	http.RLock()
	defer http.RUnlock()

	return http.Tls
}

func (http *Http) GetOrigins() []string {
	http.RLock()
	defer http.RUnlock()

	return http.Origins
}

func (http *Http) GetHeaders() []string {
	http.RLock()
	defer http.RUnlock()

	return http.Headers
}

func (http *Http) GetPort() string {
	http.RLock()
	defer http.RUnlock()

	return http.Port
}

func (http *Http) GetMode() string {
	http.RLock()
	defer http.RUnlock()

	return http.Mode
}

func (http *Http) SetHandler(name string, handler func(http go_saas_http.Http) error) {
	handlers := http.GetHandlers()

	http.Lock()
	defer http.Unlock()

	handlers[name] = handler
}
