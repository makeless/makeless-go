package makeless_go_http_basic

func (http *Http) Start() error {
	router := http.GetRouter()
	router.Use(http.CorsMiddleware(http.GetOrigins(), http.GetHeaders()))

	for _, handler := range http.GetHandlers() {
		if err := handler(http); err != nil {
			return err
		}
	}

	if http.GetTls() != nil {
		return router.RunTLS(":"+http.GetPort(), http.GetTls().GetCertPath(), http.GetTls().GetKeyPath())
	}

	return router.Run(":" + http.GetPort())
}
