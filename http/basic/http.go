package makeless_go_http_basic

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/authenticator"
	"github.com/makeless/makeless-go/database"
	"github.com/makeless/makeless-go/event"
	"github.com/makeless/makeless-go/http"
	"github.com/makeless/makeless-go/logger"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/security"
	"github.com/makeless/makeless-go/tls"
	"sync"
)

type Http struct {
	Router        *gin.Engine
	Handlers      map[string]func(http makeless_go_http.Http) error
	Logger        makeless_go_logger.Logger
	Event         makeless_go_event.Event
	Authenticator makeless_go_authenticator.Authenticator
	Security      makeless_go_security.Security
	Database      makeless_go_database.Database
	Mailer        makeless_go_mailer.Mailer
	Tls           makeless_go_tls.Tls
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

func (http *Http) GetHandlers() map[string]func(http makeless_go_http.Http) error {
	http.RLock()
	defer http.RUnlock()

	return http.Handlers
}

func (http *Http) GetLogger() makeless_go_logger.Logger {
	http.RLock()
	defer http.RUnlock()

	return http.Logger
}

func (http *Http) GetEvent() makeless_go_event.Event {
	http.RLock()
	defer http.RUnlock()

	return http.Event
}

func (http *Http) GetAuthenticator() makeless_go_authenticator.Authenticator {
	http.RLock()
	defer http.RUnlock()

	return http.Authenticator
}

func (http *Http) GetSecurity() makeless_go_security.Security {
	http.RLock()
	defer http.RUnlock()

	return http.Security
}

func (http *Http) GetDatabase() makeless_go_database.Database {
	http.RLock()
	defer http.RUnlock()

	return http.Database
}

func (http *Http) GetMailer() makeless_go_mailer.Mailer {
	http.RLock()
	defer http.RUnlock()

	return http.Mailer
}

func (http *Http) GetTls() makeless_go_tls.Tls {
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

func (http *Http) SetHandler(name string, handler func(http makeless_go_http.Http) error) {
	handlers := http.GetHandlers()

	http.Lock()
	defer http.Unlock()

	handlers[name] = handler
}
