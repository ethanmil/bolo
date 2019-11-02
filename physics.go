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

func (v *vector) getAngle() float64 {
	return math.Atan2(v.y, v.x)
}

type angle struct {
	radians float64
}

func newAngle(radians float64) angle {
	return angle{
		radians: radians,
	}
}

func (a *angle) reset() {
	a.radians = 0
}

func (a *angle) getVector() (v vector) {
	v.x = (math.Cos(a.radians))
	v.y = (math.Sin(a.radians))
	return v
}

func (a *angle) getDegrees() float64 {
	return getDegrees(a.radians)
}

func getDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

func getRadians(degrees float64) float64 {
	return degrees / (180 / math.Pi)
}
