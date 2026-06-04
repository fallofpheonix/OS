module github.com/fallofpheonix/phoenix/foundation/runtime

go 1.26

replace (
	github.com/fallofpheonix/phoenix/foundation/contracts => ../contracts
	github.com/fallofpheonix/phoenix/foundation/events => ../events
	github.com/fallofpheonix/phoenix/foundation/ledger => ../ledger
	github.com/fallofpheonix/phoenix/assurance/security => ../../assurance/security
	github.com/fallofpheonix/phoenix/foundation/distributed => ../distributed
	github.com/fallofpheonix/phoenix/foundation/runtime/kernel => ./kernel
	github.com/fallofpheonix/phoenix/governance/truth => ../../governance/truth
	github.com/fallofpheonix/phoenix/assurance/validation => ../../assurance/validation
	github.com/fallofpheonix/phoenix/foundation/observability => ../observability
	github.com/fallofpheonix/phoenix/cognition/mind => ../../cognition/mind
	github.com/fallofpheonix/phoenix/platform/os => ../../platform/os
)

require (
	github.com/fallofpheonix/phoenix/assurance/security v0.0.0
	github.com/fallofpheonix/phoenix/foundation/observability v0.0.0
	github.com/fallofpheonix/phoenix/assurance/validation v0.0.0
	google.golang.org/protobuf v1.36.11
)
