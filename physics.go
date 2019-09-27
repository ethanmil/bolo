package main

import "math"

type vector struct {
	x float64
	y float64
}

func (v *vector) reset() {
	v.x = 0
	v.y = 0
}

func (v *vector) getAngleInDegrees() float64 {
	return math.Atan2(v.y, v.x) * (180 / math.Pi)
}
