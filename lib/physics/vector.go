package physics

import "math"

// Vector -
type Vector struct {
	X float32
	Y float32
}

// NewVector -
func NewVector(x, y float32) Vector {
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
func (v *Vector) GetAngle() float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)))
}
