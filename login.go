package go_saas

import "github.com/go-saas/go-saas/http"

func (saas *Saas) login(http go_saas_http.Http) error {
	http.GetRouter().POST(
		"/api/login",
		http.GetAuthenticator().GetMiddleware().LoginHandler,
	)

	return nil
}
