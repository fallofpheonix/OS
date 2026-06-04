module github.com/fallofpheonix/phoenix/assurance/security

go 1.26

replace (
	github.com/fallofpheonix/phoenix/foundation/runtime => ../../foundation/runtime
	github.com/fallofpheonix/phoenix/foundation/runtime/kernel => ../../foundation/runtime/kernel
)

require (
	github.com/fallofpheonix/phoenix/foundation/runtime v0.0.0
	github.com/fallofpheonix/phoenix/foundation/runtime/kernel v0.0.0
)

require (
	github.com/cilium/ebpf v0.21.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
)
