module github.com/fallofpheonix/phoenix/governance/truth

go 1.26

replace github.com/fallofpheonix/phoenix/foundation/runtime => ../../internal/runtime

require (
	github.com/fallofpheonix/phoenix/foundation/runtime v0.0.0
	google.golang.org/protobuf v1.36.11
)
