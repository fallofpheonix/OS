package main

import (
	"fmt"
)

type PIDController struct {
	Kp, Ki, Kd float64
	lastError  float64
	integral   float64
}

func NewPID(kp, ki, kd float64) *PIDController {
	return &PIDController{Kp: kp, Ki: ki, Kd: kd}
}

func (c *PIDController) Update(setpoint, measured float64, dt float64) float64 {
	error := setpoint - measured
	c.integral += error * dt
	derivative := (error - c.lastError) / dt
	output := c.Kp*error + c.Ki*c.integral + c.Kd*derivative
	c.lastError = error
	return output
}

func main() {
	// Boot check
	pid := NewPID(1.0, 0.1, 0.05)
	fmt.Printf("Initial Correction: %f\n", pid.Update(100.0, 50.0, 1.0))
}
