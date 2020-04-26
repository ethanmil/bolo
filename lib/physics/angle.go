package physics

import "math"

// Angle -
type Angle float64

// NewAngle -
func NewAngle(radians float64) Angle {
	return Angle(radians)
}

// Reset -
func (a Angle) Reset() {
	a = Angle(0)
}

// GetVector -
func (a Angle) GetVector() (v Vector) {
	v.X = math.Cos(float64(a))
	v.Y = math.Sin(float64(a))
	return v
}

// GetDegrees -
func (a Angle) GetDegrees() float64 {
	return GetDegrees(float64(a))
}
