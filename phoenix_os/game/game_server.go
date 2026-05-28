package game

import (
	"fmt"
	"net/http"
	"github.com/fallofpheonix/PheonixGuard"
	ledger "github.com/fallofpheonix/PheonixTruth/src"
)

type ScoreState struct{}

func NewScoreState() *ScoreState { return &ScoreState{} }

type GameServer struct {
	Score  *ScoreState
	Warden *warden.Warden
	Ledger *ledger.Ledger
	Events string
}

func NewGameServer(s *ScoreState, w *warden.Warden, l *ledger.Ledger, e string) *GameServer {
	return &GameServer{Score: s, Warden: w, Ledger: l, Events: e}
}

func (gs *GameServer) Start(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "OK") })
	go http.ListenAndServe(addr, mux)
}
