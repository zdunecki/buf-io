package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zdunecki/buf-io/integrations/slack"
	"net/url"
	"os"
)

func host() string {
	dockerAppUrl := os.Getenv("DOCKER_APP_URL")

	if dockerAppUrl == "" {
		dockerAppUrl = "http://0.0.0.0:5555"
	}
	u, err := url.Parse(dockerAppUrl)
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
