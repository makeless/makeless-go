package makeless_go_tls

import "github.com/makeless/makeless-go/http"

type Tls interface {
	Run(http makeless_go_http.Http) error
}
