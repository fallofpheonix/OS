
go 1.26

replace (
	github.com/fallofpheonix/phoenix/assurance/security => ../assurance/security
	github.com/fallofpheonix/phoenix/foundation/runtime => ../internal/runtime
)

require github.com/fallofpheonix/phoenix/assurance/security v0.0.0

require github.com/fallofpheonix/phoenix/foundation/runtime v0.0.0 // indirect
