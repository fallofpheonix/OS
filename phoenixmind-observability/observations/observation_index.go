package observations

// Index manages the list of stored observations.
type Index struct {
    Cycles []string `json:"cycles"`
}

// AddToIndex adds a new cycle to the index.
func AddToIndex(cycle string) error {
    // TODO: Implement index update
    return nil
}
