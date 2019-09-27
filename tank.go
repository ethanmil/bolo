package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	tankSpeed = 0.1
)

type tank struct {
	element  *element
	velocity vector
}

func newTank() (t tank) {
	t.element = &element{
		size: vector{
			x: 32,
			y: 32,
		},
		chunk: sdl.Rect{
			X: 0,
			Y: 684,
			H: 32,
			W: 32,
		},
		active: true,
	}

	return t
}

func (t *tank) update() {

	keys := sdl.GetKeyboardState()

	movement := false

	t.velocity.reset()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		movement = true
		t.velocity.x = -tankSpeed
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		movement = true
		t.velocity.x = tankSpeed
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		movement = true
		t.velocity.y = tankSpeed
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		movement = true
		t.velocity.y = -tankSpeed
	}

	if movement {
		t.element.angle = t.velocity.getAngle()
	}

	t.element.position.x += t.velocity.x * delta
	t.element.position.y += t.velocity.y * delta
}
