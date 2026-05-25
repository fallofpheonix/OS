package guards
import "github.com/fallofpheonix/phoenix-os/phoenixmind-security/policies"

type ProposalGuard struct{}
func (g *ProposalGuard) CanAccess(phase string) policies.Result {
    if phase == "F1" { return policies.Block }
    return policies.Allow
}
