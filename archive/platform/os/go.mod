module github.com/fallofpheonix/phoenix/platform/os

go 1.26

replace (
	github.com/fallofpheonix/phoenix-os/phoenix_os/containment => ./phoenix_os/containment
	github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file => ./phoenix_os/containment/file
	github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network => ./phoenix_os/containment/network
	github.com/fallofpheonix/phoenix-os/phoenixmind-core => ./phoenixmind-core
	github.com/fallofpheonix/phoenix-os/phoenixmind-observability => ./observability
	github.com/fallofpheonix/phoenix/assurance/security => ../../assurance/security
	github.com/fallofpheonix/phoenix/assurance/validation => ../../assurance/validation
	github.com/fallofpheonix/phoenix/foundation/runtime => ../../internal/runtime
	github.com/fallofpheonix/phoenix/foundation/runtime/kernel => ../../internal/runtime/kernel
	github.com/fallofpheonix/phoenix/governance/truth => ../../governance/truth
	github.com/fallofpheonix/phoenixmind-validator/truth/evidence => ./phoenixmind-validator/truth/evidence
	phoenix_os/containment => ./phoenix_os/containment
)

require (
	github.com/fallofpheonix/phoenix/assurance/security v0.0.0
	github.com/fallofpheonix/phoenix/assurance/validation v0.0.0
	github.com/fallofpheonix/phoenix/foundation/runtime v0.0.0
	github.com/fallofpheonix/phoenixmind-validator/truth/evidence v0.0.0-00010101000000-000000000000
)

require (
	github.com/cilium/ebpf v0.21.0 // indirect
	github.com/fallofpheonix/phoenix/foundation/runtime/kernel v0.0.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
)
