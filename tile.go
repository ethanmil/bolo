package main

import (
	"github.com/ethanmil/go-engine/animation"
	"github.com/ethanmil/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
)

type tile struct {
	Sprite   animation.Sprite
	position physics.Vector
}

func newTile(typ string, position physics.Vector) (t tile) {
	switch typ {
	case "0": // ocean
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 0, Y: 0, H: 32, W: 32}}
		break
	case "1": // water
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 32, Y: 0, H: 32, W: 32}}
		break
	case "2": // right-bottom-road
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 64, Y: 0, H: 32, W: 32}}
		break
	case "32": // top-bottom-road
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 32, Y: 32, H: 32, W: 32}}
		break
	case "31": // right-left-road
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 0, Y: 32, H: 32, W: 32}}
		break
	case "33": // grass
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 64, Y: 32, H: 32, W: 32}}
		break
	case "34": // forest
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 96, Y: 32, H: 32, W: 32}}
		break
	case "37": // single-wall
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 192, Y: 32, H: 32, W: 32}}
		break
	case "314": // top-left-road-mine
		t.Sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 128, Y: 320, H: 32, W: 32}}
		break
	}

	t.position = position

	return t
}

func (t *tile) draw() {
	t.Sprite.Draw(t.position, 0, art, renderer)
}
