package makeless_go_authenticator_basic

import (
	"github.com/appleboy/gin-jwt/v2"
	"time"
)

func (authenticator *Authenticator) CreateMiddleware() error {
	middlware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           authenticator.GetRealm(),
		Key:             authenticator.GetKey(),
		Timeout:         authenticator.GetTimeout(),
		MaxRefresh:      authenticator.GetMaxRefresh(),
		IdentityKey:     authenticator.GetIdentityKey(),
		PayloadFunc:     authenticator.PayloadHandler,
		IdentityHandler: authenticator.IdentityHandler,
		Authenticator:   authenticator.AuthenticatorHandler,
		Authorizator:    authenticator.AuthorizatorHandler,
		Unauthorized:    authenticator.UnauthorizedHandler,
		TimeFunc:        time.Now,
		SendCookie:      true,
		SecureCookie:    authenticator.GetSecureCookie(),
		CookieHTTPOnly:  true,
		CookieName:      "jwt",
		TokenLookup:     "cookie:jwt",
		CookieDomain:    authenticator.GetCookieDomain(),
		CookieSameSite:  authenticator.GetCookieSameSite(),
	})

	if err != nil {
		return err
	}

	authenticator.SetMiddleware(middlware)
	return nil
}
