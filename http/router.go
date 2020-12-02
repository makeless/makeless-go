package makeless_go_http

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Init(http Http) error
	GetEngine() *gin.Engine
	SetEngine(engine *gin.Engine)
}
