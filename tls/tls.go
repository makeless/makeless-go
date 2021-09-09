package makeless_go_tls

import (
	"github.com/gin-gonic/gin"
)

type Tls interface {
	Run(port string, engine *gin.Engine) error
}
