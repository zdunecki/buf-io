package main

import (
	"flag"
	"github.com/zdunecki/buf-io/client"
	"github.com/zdunecki/buf-io/server"
	"strings"
)

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()

	if strings.ToLower(*flagMode) == "client" {
		client.StartClientMode()
	} else {
		server.StartServerMode()
	}
}
