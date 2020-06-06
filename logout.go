package go_saas

import "github.com/go-saas/go-saas/http"

func (saas *Saas) logout(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/logout",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.GetAuthenticator().GetMiddleware().LogoutHandler,
	)

	return nil
}
