package makeless_go_http_basic

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
	"sync"
)

type Router struct {
	engine *gin.Engine
	*sync.RWMutex
}

func (router *Router) Init(http makeless_go_http.Http) error {
	r := gin.Default()
	r.Use(http.CorsMiddleware(http.GetOrigins(), http.GetOriginsFunc(), http.GetHeaders()))

	router.SetEngine(r)
	return nil
}

func (router *Router) GetEngine() *gin.Engine {
	router.RLock()
	defer router.RUnlock()

	return router.engine
}

func (router *Router) SetEngine(engine *gin.Engine) {
	router.Lock()
	defer router.Unlock()

	router.engine = engine
}
