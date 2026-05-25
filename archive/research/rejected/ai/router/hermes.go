package router

type AgentRouter struct {
	ActiveModel string
}

func (r *AgentRouter) Route(task string) string {
	// Simple deterministic routing
	if task == "code" {
		return "qwen2.5-coder:1.5b"
	}
	return "phoenix-mind"
}
