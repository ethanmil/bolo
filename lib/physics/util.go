package physics

import "math"

// GetDegrees converts radians into degrees
func GetDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

// GetRadians converts degrees into radians
func GetRadians(degrees float64) float64 {
	return degrees / (180 / math.Pi)
}
