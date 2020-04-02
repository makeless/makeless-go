package saas_api

import "github.com/gin-gonic/gin"

func (api *Api) Extend(handler func(engine *gin.Engine)) {
	api.Lock()
	defer api.Unlock()

	api.handlers = append(api.handlers, handler)
}
