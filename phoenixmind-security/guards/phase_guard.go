package guards
import "github.com/fallofpheonix/phoenix-os/phoenixmind-security/policies"

type PhaseGuard struct{}

func (g *PhaseGuard) CanTransition(from, to string) policies.Result {
    if from == "F1" && to == "training" {
        return policies.Block
    }
    if from == "F1" && to == "proposal" {
        return policies.Block
    }
    return policies.Allow
}
