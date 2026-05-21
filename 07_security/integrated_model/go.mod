module phoenix/security/integrated_model

go 1.25.0

replace phoenix/security/physics => ../physics

replace phoenix/telemetry/entropy_engine => ../../09_telemetry/entropy_engine

replace phoenix/telemetry/process_graphs => ../../09_telemetry/process_graphs

require (
	phoenix/security/physics v0.0.0-00010101000000-000000000000
	phoenix/telemetry/entropy_engine v0.0.0-00010101000000-000000000000
	phoenix/telemetry/process_graphs v0.0.0-00010101000000-000000000000
)
