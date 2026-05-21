package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type State string

const (
	StateIdea       State = "IDEA"
	StateRFC        State = "RFC"
	StateResearch   State = "RESEARCH"
	StateExperiment State = "EXPERIMENT"
	StateBuild      State = "BUILD"
	StateRun        State = "RUN"
	StateTest       State = "TEST"
	StateDebug      State = "DEBUG"
	StateBenchmark  State = "BENCHMARK"
	StateReplay     State = "REPLAY"
	StateValidate   State = "VALIDATE"
	StateMerge      State = "MERGE"
)

var stateOrder = []State{
	StateIdea, StateRFC, StateResearch, StateExperiment,
	StateBuild, StateRun, StateTest, StateDebug,
	StateBenchmark, StateReplay, StateValidate, StateMerge,
}

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	State State  `json:"state"`
}

type Engine struct {
	mu    sync.RWMutex
	Tasks map[string]*Task
	File  string
}

func NewEngine(file string) *Engine {
	e := &Engine{
		Tasks: make(map[string]*Task),
		File:  file,
	}
	e.load()
	return e
}

func (e *Engine) load() {
	data, err := os.ReadFile(e.File)
	if err == nil {
		json.Unmarshal(data, &e.Tasks)
	}
}

func (e *Engine) save() {
	data, _ := json.MarshalIndent(e.Tasks, "", "  ")
	os.WriteFile(e.File, data, 0644)
}

func (e *Engine) CreateTask(id, title string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Tasks[id] = &Task{ID: id, Title: title, State: StateIdea}
	fmt.Printf("Task Created: %s (%s)\n", id, title)
	e.save()
}

func (e *Engine) Transition(id string, newState State) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	task, ok := e.Tasks[id]
	if !ok {
		return fmt.Errorf("task %s not found", id)
	}
	
	// Basic validation: allow any transition for now but log it
	fmt.Printf("Transition: %s [%s -> %s]\n", id, task.State, newState)
	task.State = newState
	e.save()
	return nil
}

func main() {
	engine := NewEngine("state.json")

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			engine.mu.RLock()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(engine.Tasks)
			engine.mu.RUnlock()
		} else if r.Method == http.MethodPost {
			var t Task
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			engine.CreateTask(t.ID, t.Title)
			w.WriteHeader(http.StatusCreated)
		}
	})

	http.HandleFunc("/transition", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req struct {
				ID    string `json:"id"`
				State State  `json:"state"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := engine.Transition(req.ID, req.State); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Control Plane Engine listening on :%s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
