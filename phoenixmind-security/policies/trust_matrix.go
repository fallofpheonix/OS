package policies

type Result string

const (
	Allow    Result = "ALLOW"
	Block    Result = "BLOCK"
	Escalate Result = "ESCALATE"
)

type PolicyEngine struct{}

func (e *PolicyEngine) Check(actor string, action string) Result {
	// Rule: training -> production_write: false
	if actor == "training" && action == "production_write" {
		return Block
	}
	// Rule: sandbox -> merge: false
	if actor == "sandbox" && action == "merge" {
		return Block
	}
	// Rule: agents -> kernel_access: false
	if actor == "agents" && action == "kernel_access" {
		return Block
	}
	// Rule: cognition -> runtime_access: false
	if actor == "cognition" && action == "runtime_access" {
		return Block
	}
	return Allow
}
