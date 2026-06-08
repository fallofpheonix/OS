module github.com/fallofpheonix/phoenix/platform/cli

go 1.26

replace (
	github.com/fallofpheonix/phoenix/foundation/contracts => ../../foundation/contracts
	github.com/fallofpheonix/phoenix/foundation/events => ../../foundation/events
	github.com/fallofpheonix/phoenix/foundation/ledger => ../../internal/ledger
	github.com/fallofpheonix/phoenix/foundation/math => ../../internal/math
	github.com/fallofpheonix/phoenix/foundation/runtime => ../../internal/runtime
)

require (
	github.com/fallofpheonix/phoenix/foundation/runtime v0.0.0
)

require (
	github.com/fallofpheonix/phoenix/foundation/ledger v0.0.0 // indirect
	github.com/fallofpheonix/phoenix/foundation/math v0.0.0 // indirect
)
