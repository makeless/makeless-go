package makeless_go_http_basic

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/database"
	"github.com/makeless/makeless-go/http"
	h "net/http"
	"sync"
)

type Router struct {
	engine   *gin.Engine
	Database makeless_go_database.Database
	*sync.RWMutex
}

func (router *Router) Init(http makeless_go_http.Http) error {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.AbortWithStatus(h.StatusInternalServerError)
	}))
	r.Use(http.CorsMiddleware(http.GetOrigins(), http.GetHeaders()))

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

func (router *Router) GetDatabase() makeless_go_database.Database {
	router.RLock()
	defer router.RUnlock()

	return router.Database
}
