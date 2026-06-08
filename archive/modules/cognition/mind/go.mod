
go 1.26

replace (
	github.com/fallofpheonix/phoenix/foundation/runtime => ../../internal/runtime
	github.com/fallofpheonix/phoenix/foundation/ledger => ../../internal/ledger
	github.com/fallofpheonix/phoenix/assurance/security => ../../assurance/security
	github.com/fallofpheonix/phoenix/governance/truth => ../../governance/truth
	github.com/fallofpheonix/phoenix/foundation/runtime/kernel => ../../internal/runtime/kernel
	github.com/fallofpheonix/phoenix/foundation/observability => ../../foundation/observability
)

require (
	github.com/fallofpheonix/phoenix/foundation/runtime v0.0.0
	github.com/fallofpheonix/phoenix/foundation/ledger v0.0.0
	github.com/fallofpheonix/phoenix/assurance/security v0.0.0
)

require github.com/mattn/go-sqlite3 v1.14.44 // indirect
