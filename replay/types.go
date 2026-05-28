package replay

type ReplayResult struct {
    InputHash  string
    OutputHash string
    Divergence bool
    EvidenceID string
}
