package makeless_go

import (
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/http"
)

func (makeless *Makeless) events(http makeless_go_http.Http) error {
	http.GetRouter().GetEngine().GET(
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

			go func() {
				if err := http.GetEvent().Trigger(userId, http.GetEvent().GetName(), "subscribed", clientId); err != nil {
					http.GetEvent().TriggerError(err)
				}
			}()

			for {
				select {
				case event := <-http.GetEvent().Listen(userId, clientId):
					sseEvent := sse.Event{
						Event: event.GetChannel(),
						Retry: 3,
						Data:  event.GetData(),
					}

					if err := sse.Encode(w, sseEvent); err != nil {
						panic(err)
					}

					w.Flush()
				case err := <-http.GetEvent().ListenError():
					panic(err)
				case <-w.CloseNotify():
					http.GetEvent().Unsubscribe(userId, clientId)

					if err := http.GetEvent().Trigger(userId, http.GetEvent().GetName(), "unsubscribed", clientId); err != nil {
						panic(err)
					}
				}
			}
		},
	)

	return nil
}
