package memory

type Pattern struct {
	FailureSignature string
	RecommendedFix   string
	Risk             string
}

var KnownPatterns = []Pattern{
	{FailureSignature: "memory_leak", RecommendedFix: "buffer_limit", Risk: "MEDIUM"},
}
