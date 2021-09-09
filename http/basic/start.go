package makeless_go_http_basic

func (http *Http) Start() error {
	if err := http.GetRouter().Init(http); err != nil {
		return err
	}

	for _, handler := range http.GetHandlers() {
		if err := handler(http); err != nil {
			return err
		}
	}

	if http.GetTls() != nil {
		return http.GetTls().Run(http.GetPort(), http.GetRouter().GetEngine())
	}

	return http.GetRouter().GetEngine().Run(":" + http.GetPort())
}
