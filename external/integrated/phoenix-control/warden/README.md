# Phoenix Warden (Control & Actuation)

Closed-loop response engine for PhoenixOS.

## Purpose
Throttles or isolates suspicious processes using PID-based feedback control (Cgroups/Nice).

## Validation Gates
- [ ] Settling time < 2 seconds.
- [ ] Overshoot < 10%.
- [ ] Fail-open safety (never throttle init/critical sys).
