package physics

import "math"

// GetDegrees converts radians into degrees
func GetDegrees(radians float32) float32 {
	return radians * (180 / math.Pi)
}

// GetRadians converts degrees into radians
func GetRadians(degrees float32) float32 {
	return degrees / (180 / math.Pi)
}
