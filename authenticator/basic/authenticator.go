package makeless_go_authenticator_basic

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/security"
	"net/http"
	"sync"
	"time"
)

type Authenticator struct {
	Middlware      *jwt.GinJWTMiddleware
	Security       makeless_go_security.Security
	Realm          string
	Key            string
	Timeout        time.Duration
	MaxRefresh     time.Duration
	IdentityKey    string
	SecureCookie   bool
	CookieDomain   string
	CookieSameSite http.SameSite
	*sync.RWMutex
}

func (authenticator *Authenticator) SetMiddleware(middleware *jwt.GinJWTMiddleware) {
	authenticator.Lock()
	defer authenticator.Unlock()

	authenticator.Middlware = middleware
}

func (authenticator *Authenticator) GetMiddleware() *jwt.GinJWTMiddleware {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.Middlware
}

func (authenticator *Authenticator) AuthMiddleware() gin.HandlerFunc {
	return authenticator.GetMiddleware().MiddlewareFunc()
}

func (authenticator *Authenticator) GetSecurity() makeless_go_security.Security {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.Security
}

func (authenticator *Authenticator) GetRealm() string {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.Realm
}

func (authenticator *Authenticator) GetKey() []byte {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return []byte(authenticator.Key)
}

func (authenticator *Authenticator) GetTimeout() time.Duration {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.Timeout
}

func (authenticator *Authenticator) GetMaxRefresh() time.Duration {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.MaxRefresh
}

func (authenticator *Authenticator) GetIdentityKey() string {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.Key
}

func (authenticator *Authenticator) GetSecureCookie() bool {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.SecureCookie
}

func (authenticator *Authenticator) GetCookieDomain() string {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.CookieDomain
}

func (authenticator *Authenticator) GetCookieSameSite() http.SameSite {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.CookieSameSite
}
