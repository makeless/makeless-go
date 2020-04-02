package saas_api

func (api *Api) Extend(handler func(api *Api)) {
	api.Lock()
	defer api.Unlock()

	api.handlers = append(api.handlers, handler)
}
