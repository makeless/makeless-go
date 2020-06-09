package go_saas

import "github.com/go-saas/go-saas/http"

func (saas *Saas) login(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/login",
		http.GetAuthenticator().GetMiddleware().LoginHandler,
	)

	return nil
}

func (saas *Saas) logout(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/logout",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.GetAuthenticator().GetMiddleware().LogoutHandler,
	)

	return nil
}

func (saas *Saas) refreshToken(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/refresh-token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.GetAuthenticator().GetMiddleware().RefreshHandler,
	)

	return nil
}
