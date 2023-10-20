package services

import (
	"commando/models"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gorilla/websocket"
)

// TODO: expose ports as configuration or environment variables
const serverPort = "1111"
const clientPort = "1112"

type WebServer struct {
	BroadcastChannel models.BroadcastChannel
}

func (server WebServer) Start() {
	// http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", server.serveWebSockets)

	go startClient()

	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		log.Fatalln(err)
	}

}

// func serveHome(writer http.ResponseWriter, req *http.Request) {
// 	http.ServeFile(writer, req, "views/index.html")
// }

func startClient() {
	cmd := exec.Command("yarn", "dev")
	cmd.Dir = "./client"

	openBrowser("http://localhost:" + clientPort)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: look at a better solution than enabling all origins
	},
}

var clients = make(map[*websocket.Conn]bool)

func (server WebServer) serveWebSockets(writer http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()
	clients[conn] = true

	for {
		msg := <-server.BroadcastChannel
		for client, connected := range clients {
			if connected {
				err := client.WriteJSON(msg)
				if err != nil {
					delete(clients, conn)
					return
				}
			}
		}
	}
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Print(err)
	}
}
