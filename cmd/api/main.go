package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zdunecki/buf-io/integrations/slack"
	"net/url"
	"os"
)

func host() string {
	u, err := url.Parse(os.Getenv("DOCKER_APP_URL"))
	if err != nil {
		panic(err)
	}

	return u.Host
}

func main() {
	r := gin.Default()

	s := r.Group("/slack")
	{
		s.POST("/events", slack.Events)
		s.POST("/interactive-components", slack.InteractiveComponents)
	}

	if err := r.Run(host()); err != nil {
		panic(err)
	}
}
