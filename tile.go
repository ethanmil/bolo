package main

import "github.com/veandco/go-sdl2/sdl"

type tile struct {
	sprite   sprite
	position vector
}

func newTile(typ string, position vector) (t tile) {
	switch typ {
	case "0": // ocean
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 0, Y: 0, H: 32, W: 32}}
		break
	case "1": // water
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 32, Y: 0, H: 32, W: 32}}
		break
	case "2": // right-bottom-road
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 64, Y: 0, H: 32, W: 32}}
		break
	case "32": // top-bottom-road
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 32, Y: 32, H: 32, W: 32}}
		break
	case "31": // right-left-road
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 0, Y: 32, H: 32, W: 32}}
		break
	case "33": // grass
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 64, Y: 32, H: 32, W: 32}}
		break
	case "34": // forest
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 96, Y: 32, H: 32, W: 32}}
		break
	case "37": // single-wall
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 192, Y: 32, H: 32, W: 32}}
		break
	case "314": // top-left-road-mine
		t.sprite = sprite{size: vector{x: 32, y: 32}, chunk: sdl.Rect{X: 128, Y: 320, H: 32, W: 32}}
		break
	}

	t.position = position

	return t
}

func (t *tile) draw() {
	t.sprite.draw(t.position, 0, renderer)
}
