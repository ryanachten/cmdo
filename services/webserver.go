package services

import (
	"commando/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebServer struct {
	BroadcastChannel models.BroadcastChannel
}

func (ws WebServer) Start() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", ws.serveWebSockets)

	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func serveHome(writer http.ResponseWriter, req *http.Request) {
	http.ServeFile(writer, req, "views/index.html")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ws WebServer) serveWebSockets(writer http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	for {
		msg := <-ws.BroadcastChannel
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
