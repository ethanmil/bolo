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
	return math.Atan2(v.y, v.x) * (180 / math.Pi)
}

type angle struct {
	degrees float64
}

func newAngle(degrees float64) angle {
	return angle{
		degrees: degrees,
	}
}

func (a *angle) reset() {
	a.degrees = 0
}

func (a *angle) getVector() (v vector) {
	v.x = getDegrees(math.Cos(getRadians(a.degrees))) / 360 // dividing by 360 scales the velocity to 0-1
	v.y = getDegrees(math.Sin(getRadians(a.degrees))) / 360 // dividing by 360 scales the velocity to 0-1
	return v
}

func getDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

func getRadians(degrees float64) float64 {
	return degrees / (180 / math.Pi)
}
