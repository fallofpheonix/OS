package multiplayer

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/fallofpheonix/phoenix/game/engine"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for the bridge
	},
}

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	clientsMu sync.Mutex
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (s *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[WS ERROR] Upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	s.clientsMu.Lock()
	s.clients[conn] = true
	s.clientsMu.Unlock()

	log.Printf("[WS] Client connected")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			s.clientsMu.Lock()
			delete(s.clients, conn)
			s.clientsMu.Unlock()
			log.Printf("[WS] Client disconnected")
			break
		}
	}
}

func (s *WebSocketServer) BroadcastWorldState(state engine.WorldState) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	data, err := json.Marshal(state)
	if err != nil {
		log.Printf("[WS ERROR] Marshal failed: %v", err)
		return
	}

	for client := range s.clients {
		err := client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Printf("[WS ERROR] Write failed: %v", err)
			client.Close()
			delete(s.clients, client)
		}
	}
}

func (s *WebSocketServer) Start(addr string) {
	http.HandleFunc("/ws", s.HandleConnections)
	log.Printf("[WS SERVER] Listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("[WS SERVER ERROR] %v", err)
	}
}
