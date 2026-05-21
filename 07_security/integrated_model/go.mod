module sentinel/security/integrated_model

go 1.25.0

replace sentinel/security/physics => ../physics

replace sentinel/telemetry/entropy_engine => ../../09_telemetry/entropy_engine

replace sentinel/telemetry/process_graphs => ../../09_telemetry/process_graphs

require (
	sentinel/security/physics v0.0.0-00010101000000-000000000000
	sentinel/telemetry/entropy_engine v0.0.0-00010101000000-000000000000
	sentinel/telemetry/process_graphs v0.0.0-00010101000000-000000000000
)
