module github.com/fallofpheonix/phoenix/foundation/runtime

go 1.26

require (
	github.com/fallofpheonix/phoenix/internal/contracts v0.0.0
	github.com/fallofpheonix/phoenix/internal/protocol v0.0.0
	github.com/fallofpheonix/phoenix/internal/state v0.0.0
	github.com/fallofpheonix/phoenix/internal/consensus v0.0.0
	github.com/fallofpheonix/phoenix/foundation/math v0.0.0
)

replace github.com/fallofpheonix/phoenix/internal/contracts => ../contracts
replace github.com/fallofpheonix/phoenix/internal/protocol => ../protocol
replace github.com/fallofpheonix/phoenix/internal/state => ../state
replace github.com/fallofpheonix/phoenix/internal/consensus => ../consensus
replace github.com/fallofpheonix/phoenix/foundation/math => ../math


