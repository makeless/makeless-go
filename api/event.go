package saas_api

import (
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api *Api) events(c *gin.Context) {
	userId, err := api.GetUserId(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response(err.Error(), nil))
		return
	}

	w := c.Writer
	h := w.Header()
	h.Set("Content-Type", "text/event-stream")
	h.Set("Cache-Control", "no-cache")
	h.Set("Connection", "keep-alive")
	h.Set("X-Accel-Buffering", "no")

	clientId := api.GetEvent().NewClientId()
	api.GetEvent().Subscribe(userId, clientId)

	go func() {
		api.GetEvent().Emit(userId, sse.Event{
			Data: fmt.Sprintf("subscribed: %d", clientId),
		})
	}()

	for {
		select {
		case event := <-api.GetEvent().Listen(userId, clientId):
			if err := sse.Encode(w, event); err != nil {
				api.GetLogger().Println(err.Error())
			}
			w.Flush()
		case <-w.CloseNotify():
			api.GetEvent().Unsubscribe(userId, clientId)
			api.GetEvent().Emit(userId, sse.Event{
				Data: fmt.Sprintf("unsubscribed: %d", clientId),
			})
			return
		}
	}
}
