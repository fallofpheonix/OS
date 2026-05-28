package repair

type RollbackEstimator struct{}

func (r *RollbackEstimator) Estimate(change string) bool {
	return false
}
