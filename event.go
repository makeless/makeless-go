package go_saas

import (
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/http"
)

func (saas *Saas) events(http go_saas_http.Http) error {
	http.GetRouter().GET(
		"/api/auth/event",
		http.GetAuthenticator().GetMiddleware().MiddlewareFunc(),
		func(c *gin.Context) {
			userId := http.GetAuthenticator().GetAuthUserId(c)

			w := c.Writer
			h := w.Header()
			h.Set("Content-Type", "text/event-stream")
			h.Set("Cache-Control", "no-cache")
			h.Set("Connection", "keep-alive")
			h.Set("X-Accel-Buffering", "no")

			clientId := http.GetEvent().NewClientId()
			http.GetEvent().Subscribe(userId, clientId)

			go http.GetEvent().Trigger(userId, "go-saas", "subscribed", clientId)

			for {
				select {
				case event := <-http.GetEvent().Listen(userId, clientId):
					if err := sse.Encode(w, event); err != nil {
						http.GetLogger().Println(err.Error())
					}
					w.Flush()
				case <-w.CloseNotify():
					http.GetEvent().Unsubscribe(userId, clientId)
					http.GetEvent().Trigger(userId, "go-saas", "unsubscribed", clientId)
				}
			}
		},
	)

	return nil
}
