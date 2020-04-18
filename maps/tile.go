package maps

import (
	"github.com/ethanmil/go-engine/animation"
	"github.com/ethanmil/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
)

// Tile -
type Tile struct {
	sprite   animation.Sprite
	position physics.Vector
}

// NewTile -
func NewTile(typ string, position physics.Vector) (t Tile) {
	switch typ {
	case "0": // ocean
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 0, Y: 0, H: 32, W: 32}}
		break
	case "1": // water
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 32, Y: 0, H: 32, W: 32}}
		break
	case "2": // right-bottom-road
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 64, Y: 0, H: 32, W: 32}}
		break
	case "32": // top-bottom-road
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 32, Y: 32, H: 32, W: 32}}
		break
	case "31": // right-left-road
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 0, Y: 32, H: 32, W: 32}}
		break
	case "33": // grass
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 64, Y: 32, H: 32, W: 32}}
		break
	case "34": // forest
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 96, Y: 32, H: 32, W: 32}}
		break
	case "37": // single-wall
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 192, Y: 32, H: 32, W: 32}}
		break
	case "314": // top-left-road-mine
		t.sprite = animation.Sprite{Size: physics.Vector{X: 32, Y: 32}, Chunk: sdl.Rect{X: 128, Y: 320, H: 32, W: 32}}
		break
	}

	t.position = position

	return t
}

// Draw -
func (t *Tile) Draw(texture *sdl.Texture, renderer *sdl.Renderer) {
	t.sprite.Draw(t.position, 0, texture, renderer)
}
