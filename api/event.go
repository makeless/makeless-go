package saas_api

import (
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

	if err := sse.Encode(w, sse.Event{
		Data: "test",
	}); err != nil {
		api.GetLogger().Println(err.Error())
	}
	w.Flush()

	for {
		select {
		case event, ok := <-api.GetEvent().Listen(userId, clientId):
			if !ok {
				continue
			}

			if err := sse.Encode(w, event); err != nil {
				api.GetLogger().Println(err.Error())
			}
			w.Flush()
		case <-w.CloseNotify():
			api.GetEvent().Unsubscribe(userId, clientId)
			return
		}
	}
}
