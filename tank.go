package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	tankSpeed = 0.05
)

type tank struct {
	element *element
}

func newTank() (t tank) {
	t.element = &element{
		size: vector{
			x: 32,
			y: 32,
		},
		chunk: sdl.Rect{
			X: 130,
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

	switch uint8(1) {
	case keys[sdl.SCANCODE_LEFT]:
		t.element.angle = 270
		t.element.position.x -= tankSpeed
		break
	case keys[sdl.SCANCODE_RIGHT]:
		t.element.angle = 90
		t.element.position.x += tankSpeed
		break
	}
}
