package slack

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Events(c *gin.Context) {
	event := &Event{}

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(data, &event); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if event.Type != "event_callback" {
		c.JSON(http.StatusOK, event.Challenge)
		return
	}

	var eventError error

	switch event.Event.Type {
	case string(EventType.FileCreated):
		eventError = EventType.FileCreated.callback(data, event)
	default:
		{
			c.JSON(http.StatusNotFound, nil)
			return
		}
	}

	if eventError != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"ok": true})
}

func InteractiveComponents(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	payload, err := url.QueryUnescape(string(body[8:]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var msg InteractiveComponent
	if err = json.Unmarshal([]byte(payload), &msg); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	switch msg.Type {
	case string(InteractiveComponentType.InteractiveMessage):
		InteractiveComponentType.InteractiveMessage.callback(msg)
	default:
		{
			c.JSON(http.StatusNotFound, nil)
			return
		}
	}

	var val InteractiveMessageValueCallback
	if err := json.Unmarshal([]byte(msg.Actions[0].Value), &val); err != nil {
		log.Fatal(err)
	}

	if val.Val == "1" {
		c.JSON(http.StatusOK, map[string]interface{}{"text": "We've started uploading a file"})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"text": "Upload gets cancelled"})
	}
}
