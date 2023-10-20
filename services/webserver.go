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

type WebServer struct {
	BroadcastChannel models.BroadcastChannel
}

func (ws WebServer) Start() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", ws.serveWebSockets)

	openBrowser("http://localhost:1111")

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
