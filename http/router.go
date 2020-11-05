package makeless_go_http

import (
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/database"
)

type Router interface {
	Init(http Http) error
	GetEngine() *gin.Engine
	SetEngine(engine *gin.Engine)
	GetDatabase() makeless_go_database.Database
}
