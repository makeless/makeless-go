package makeless_go

import "github.com/makeless/makeless-go/http"

func (makeless *Makeless) login(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().POST(
		"/api/login",
		http.GetAuthenticator().GetMiddleware().LoginHandler,
	)

	return nil
}

func (makeless *Makeless) logout(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().GET(
		"/api/auth/logout",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.GetAuthenticator().GetMiddleware().LogoutHandler,
	)

	return nil
}

func (makeless *Makeless) refreshToken(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().GET(
		"/api/auth/refresh-token",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		http.GetAuthenticator().GetMiddleware().RefreshHandler,
	)

	return nil
}
