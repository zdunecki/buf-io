package server

import (
	"fmt"
	"github.com/zdunecki/buf-io/integrations/dropbox"
	"github.com/zdunecki/buf-io/integrations/slack"
	"github.com/zdunecki/buf-io/sockets/client"
	"github.com/zdunecki/buf-io/sockets/services/file"
	"net"
)

type ClientManager struct {
	clients       map[*client.Client]bool
	broadcast     chan []byte
	register      chan *client.Client
	unregister    chan *client.Client
	currentClient *client.Client
}

const BufferSize int = 4096

func (manager *ClientManager) send(client *client.Client) {
	defer client.Socket.Close()
	for {
		select {
		case message, ok := <-client.Data:
			if !ok {
				return
			}
			_, _ = client.Socket.Write(message)
		}
	}
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.Data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.broadcast:
			for connection := range manager.clients {
				if connection == manager.currentClient {
					select {
					case connection.Data <- message:
					default:
						close(connection.Data)
						delete(manager.clients, connection)
						manager.currentClient = nil
					}
				}
			}
		}
	}
}

func (manager *ClientManager) receive(client *client.Client) {
	for {
		bufferFileName := make([]byte, 64)
		bufferFileSize := make([]byte, 10)

		_, err := client.Socket.Read(bufferFileName)
		if err != nil {
			manager.unregister <- client
			client.Socket.Close()
			break
		}

		_, err = client.Socket.Read(bufferFileSize)
		if err != nil {
			panic(err)
		}

		fileName := file.GetName(bufferFileName)
		fileSize, err := file.GetSize(bufferFileSize)
		if err != nil {
			panic(err)
		}

		var receivedBytes int
		var receivedContent []byte

		content := make([]byte, BufferSize)

		for {
			if receivedBytes >= fileSize {
				break
			}
			contentLeft := fileSize - receivedBytes
			if contentLeft < BufferSize {
				content = make([]byte, contentLeft)
			}

			_, err := client.Socket.Read(content)
			if err != nil {
				panic(err)
			}

			receivedContent = append(receivedContent, content...)
			receivedBytes += BufferSize
		}

		go dropbox.Upload(fileName, receivedContent)
		go slack.Upload(fileName, receivedContent)

		manager.broadcast <- []byte("done: " + fileName + "\n")
		manager.currentClient = client
	}
}

func StartServerMode() {
	fmt.Println("Starting server...")
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println(err)
	}
	manager := ClientManager{
		clients:    make(map[*client.Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *client.Client),
		unregister: make(chan *client.Client),
	}
	go manager.start()
	for {
		connection, err := listener.Accept()

		if err != nil {
			panic(err)
		}
		defer connection.Close()

		c := &client.Client{Socket: connection, Data: make(chan []byte)}
		manager.register <- c
		go manager.receive(c)
		go manager.send(c)
	}
}
