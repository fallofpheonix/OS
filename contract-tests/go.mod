module github.com/fallofpheonix/phoenix/contract-tests

go 1.26

replace (
	github.com/fallofpheonix/phoenix/foundation/runtime => ../foundation/runtime
	github.com/fallofpheonix/phoenix/assurance/security => ../assurance/security
)

require (
	github.com/fallofpheonix/phoenix/foundation/runtime v0.0.0
	github.com/fallofpheonix/phoenix/assurance/security v0.0.0
)
