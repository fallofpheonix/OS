package nexus
func ValidateBFTQuorum(votes int, totalNodes int) bool { return votes > (totalNodes * 2 / 3) }
