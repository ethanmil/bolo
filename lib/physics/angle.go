package physics

import "math"

// Angle -
type Angle float32

// NewAngle -
func NewAngle(radians float32) Angle {
	return Angle(radians)
}

// Reset -
func (a Angle) Reset() {
	a = Angle(0)
}

// GetVector -
func (a Angle) GetVector() (v Vector) {
	v.X = float32(math.Cos(float64(a)))
	v.Y = float32(math.Sin(float64(a)))
	return v
}

// GetDegrees -
func (a Angle) GetDegrees() float32 {
	return GetDegrees(float32(a))
}
