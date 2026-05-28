package sandbox

type ApprovalGate struct{}

func (g *ApprovalGate) NeedsHuman(risk string) bool {
    return risk == "DANGEROUS"
}
