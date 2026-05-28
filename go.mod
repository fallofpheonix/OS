module github.com/fallofpheonix/phoenix-os

go 1.26

replace github.com/fallofpheonix/phoenix-os/phoenix_os/containment => ./phoenix_os/containment

replace github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file => ./phoenix_os/containment/file

replace github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network => ./phoenix_os/containment/network

replace phoenix_os/containment => ./phoenix_os/containment

replace github.com/fallofpheonix/phoenix-os/phoenixmind-core => ./core/pheonixmind-core

replace github.com/fallofpheonix/PheonixDistributed => ../PheonixDistributed

replace github.com/fallofpheonix/PheonixSimulation => ../PheonixSimulation

replace github.com/fallofpheonix/PheonixTruth => ../PheonixTruth

replace github.com/fallofpheonix/PheonixGuard => ../PheonixGuard

require github.com/cilium/ebpf v0.21.0 // indirect

require github.com/fallofpheonix/PheonixDistributed v0.0.0

require github.com/fallofpheonix/PheonixTruth v0.0.0

require github.com/fallofpheonix/PheonixGuard v0.0.0

require (
	github.com/fallofpheonix/PheonixKernel v0.0.0-00010101000000-000000000000
	github.com/fallofpheonix/phoenix-os/phoenixmind-core v0.0.0
	github.com/fallofpheonix/phoenixmind-validator/truth/evidence v0.0.0-00010101000000-000000000000
)

require golang.org/x/sys v0.37.0 // indirect

replace github.com/fallofpheonix/phoenixmind-validator/truth/evidence => ./phoenixmind-validator/truth/evidence

replace github.com/fallofpheonix/PheonixKernel => ../PheonixKernel
