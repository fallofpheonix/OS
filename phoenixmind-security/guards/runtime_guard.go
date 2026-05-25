package guards
import "github.com/fallofpheonix/phoenix-os/phoenixmind-security/policies"

type RuntimeGuard struct {
	Engine *policies.PolicyEngine
}

func (g *RuntimeGuard) CheckAction(actor string, action string) policies.Result {
	if actor == "cognition" && action == "runtime_write" {
		return policies.Block
	}
	if actor == "runtime" && action == "patch" {
		return policies.Block
	}
	return g.Engine.Check(actor, action)
}
