package multiplayer

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix/game/engine"
	"github.com/gorilla/websocket"
)

func TestWebSocketServer_Integration(t *testing.T) {
	s := NewWebSocketServer()
	server := httptest.NewServer(http.HandlerFunc(s.HandleConnections))
	defer server.Close()

	// Convert http URL to ws URL
	url := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

	// Connect mock client
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
	defer conn.Close()

	// Broadcast a test state
	ws := engine.NewWorldState(0)
	ws.Tick = 42
	engine.InjectEntity(ws, &engine.Entity{ID: "test", Status: "TESTING"})

	// Wait a bit for connection to be registered
	time.Sleep(100 * time.Millisecond)
	s.BroadcastWorldState(*ws)

	// Read message from client
	_, message, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("ReadMessage failed: %v", err)
	}

	if !strings.Contains(string(message), "\"tick\":42") {
		t.Errorf("Expected tick 42 in message, got %s", string(message))
	}
}
