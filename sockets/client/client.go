package client

import (
	"bufio"
	"fmt"
	"github.com/zdunecki/buf-io/sockets/services/file"
	"net"
	"os"
	"strings"
)

type Client struct {
	Socket net.Conn
	Data   chan []byte
}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.Socket.Read(message)
		if err != nil {
			client.Socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
		}
	}
}

func m(a interface{}, e interface{}) interface{} {
	switch e.(type) {
	case error:
		return e
	default:
		return a
	}
}

func mo(result ...interface{}) (interface{}, error) {
	var f []interface{}

	for _, r := range result {
		switch r.(type) {
		case error:
			return nil, r.(error)
		default:
			f = append(f, r)
		}
	}

	return f, nil
}

func StartClientMode() {
	fmt.Println("Starting client...")
	connection, error := net.Dial("tcp", "localhost:12345")
	defer connection.Close()

	if error != nil {
		fmt.Println(error)
	}
	client := &Client{Socket: connection}
	go client.receive()
	for {
		reader := bufio.NewReader(os.Stdin)
		path, _ := reader.ReadString('\n')
		path = strings.TrimRight(path, "\n")

		for _, f := range file.CreateFormat(path) {
			if _, err := mo(
				m(connection.Write(f.Name)),
				m(connection.Write(f.Size)),
				m(connection.Write(f.Content)),
			); err != nil {
				panic(err)
			}
		}
	}
}
