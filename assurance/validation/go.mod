module github.com/fallofpheonix/phoenix/assurance/validation

go 1.26

replace (
	github.com/fallofpheonix/phoenix/foundation/contracts => ../../foundation/contracts
	github.com/fallofpheonix/phoenix/foundation/events => ../../foundation/events
	github.com/fallofpheonix/phoenix/foundation/ledger => ../../foundation/ledger
	github.com/fallofpheonix/phoenix/foundation/runtime => ../../foundation/runtime
	github.com/fallofpheonix/phoenix/assurance/security => ../security
	github.com/fallofpheonix/phoenix/foundation/distributed => ../../foundation/distributed
	github.com/fallofpheonix/phoenix/governance/truth => ../../governance/truth
	github.com/fallofpheonix/phoenix/foundation/runtime/kernel => ../../foundation/runtime/kernel
	github.com/fallofpheonix/phoenix/foundation/observability => ../../foundation/observability
	github.com/fallofpheonix/phoenix/cognition => ../../cognition
	github.com/fallofpheonix/phoenix/cognition/mind => ../../cognition/mind
	github.com/fallofpheonix/phoenix/platform/os => ../../platform/os
)

require (
	github.com/fallofpheonix/phoenix/foundation/contracts v0.0.0
	github.com/fallofpheonix/phoenix/foundation/events v0.0.0
	github.com/fallofpheonix/phoenix/foundation/ledger v0.0.0
	github.com/fallofpheonix/phoenix/foundation/runtime v0.0.0
	github.com/fallofpheonix/phoenix/assurance/security v0.0.0
	github.com/fallofpheonix/phoenix/foundation/runtime/kernel v0.0.0
	github.com/fallofpheonix/phoenix/cognition/mind v0.0.0
	github.com/fallofpheonix/phoenix/foundation/observability v0.0.0
	github.com/fallofpheonix/phoenix/platform/os v0.0.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/fallofpheonix/phoenix/cognition v0.0.0-00010101000000-000000000000 // indirect
	github.com/fallofpheonix/phoenix/foundation/distributed v0.0.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.44 // indirect
	golang.org/x/sys v0.37.0 // indirect
)
