package makeless_go_http_basic

func (http *Http) Start() error {
	for _, handler := range http.GetHandlers() {
		if handler == nil {
			continue
		}

		if err := handler(http); err != nil {
			return err
		}
	}

	if http.GetTls() != nil {
		return http.GetTls().Run(http.GetPort(), http.GetRouter().GetEngine())
	}

	return http.GetRouter().GetEngine().Run(":" + http.GetPort())
}
