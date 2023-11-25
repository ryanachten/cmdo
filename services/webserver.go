package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gorilla/websocket"
	"github.com/ryanachten/cmdo/events"
)

// TODO: expose ports as configuration or environment variables
const serverPort = "1111"

var history = make([]events.BroadcastMessage, 0)
var currentState = make(map[string]events.CommandStateType)

type WebServer struct {
	EventBus events.EventBus
}

func (server WebServer) Start() {
	http.HandleFunc("/", serveHome)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./views/static"))))

	http.HandleFunc("/api/state", server.serveState)
	http.HandleFunc("/api/history", server.serveHistory)
	http.HandleFunc("/api/command", server.handleCommandRequest)

	http.HandleFunc("/ws", server.serveWebSockets)

	openBrowser("http://localhost:" + serverPort)

	err := http.ListenAndServe("127.0.0.1:"+serverPort, nil)
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

var clients = make(map[*websocket.Conn]bool)

func (server WebServer) serveWebSockets(writer http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()
	clients[conn] = true

	for {
		select {
		case msg := <-server.EventBus.CommandOutput:
			history = append(history, msg)
			broadcastMessage(msg, conn)
		case msg := <-server.EventBus.CommandState:
			currentState[msg.CommandName] = msg.MessageType
			broadcastMessage(msg, conn)
		}
	}
}

func broadcastMessage(msg events.BroadcastMessage, conn *websocket.Conn) {
	for client, connected := range clients {
		if connected {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending JSON message: %v\n", err)
				delete(clients, conn)
				return
			}
		}
	}
}

func (server WebServer) serveState(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(currentState)
}

func (server WebServer) serveHistory(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(history)
}

func (server WebServer) handleCommandRequest(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	commandRequest := events.CommandRequest{}
	json.NewDecoder(req.Body).Decode(&commandRequest)
	server.EventBus.CommandRequest <- commandRequest
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
		log.Printf("Error opening browser: %v\n", err)
	}
}
