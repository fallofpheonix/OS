package guards
import "github.com/fallofpheonix/phoenix-os/phoenixmind-security/policies"

type MergeGuard struct{}
func (g *MergeGuard) CanMerge(actor string) policies.Result { 
    if actor == "observation" || actor == "sandbox" {
        return policies.Block
    }
    return policies.Allow 
}
