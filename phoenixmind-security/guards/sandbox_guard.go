package guards
import "github.com/fallofpheonix/phoenix-os/phoenixmind-security/policies"

type SandboxGuard struct{}
func (g *SandboxGuard) CanMerge() policies.Result { return policies.Block }
