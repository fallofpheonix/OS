package memory

type Strategy struct {
	Name string
	Body string
}

type StrategyStore struct {
	Strategies map[string]Strategy
}
