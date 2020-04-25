package physics

import "math"

// Vector -
type Vector struct {
	X float64
	Y float64
}

// NewVector -
func NewVector(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

// Reset -
func (v *Vector) Reset() {
	v.X = 0
	v.Y = 0
}

// GetAngle -
func (v *Vector) GetAngle() float64 {
	return math.Atan2(v.Y, v.X)
}
