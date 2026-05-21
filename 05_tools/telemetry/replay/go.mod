module phoenix/tools/replay

go 1.25.0

replace phoenix/telemetry => ../../../09_telemetry

replace phoenix/security => ../../../07_security

replace phoenix/ledger => ../../../phoenix_os/ledger

replace phoenix/security/game => ../../../07_security/game

require (
	phoenix/ai v0.0.0-00010101000000-000000000000
	phoenix/ledger v0.0.0-00010101000000-000000000000
	phoenix/security v0.0.0-00010101000000-000000000000
	phoenix/telemetry v0.0.0-00010101000000-000000000000
)

replace phoenix/ai => ../../../06_ai

replace phoenix/kernel => ../../../10_kernel
