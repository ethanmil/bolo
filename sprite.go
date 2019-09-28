package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type sprite struct {
	size  vector
	chunk sdl.Rect
	angle float64

	// log helper
	lastLogged time.Time
}

func (s *sprite) draw(position vector, angle float64, renderer *sdl.Renderer) {
	renderer.CopyEx(
		art,
		&sdl.Rect{X: int32(s.chunk.X), Y: int32(s.chunk.Y), W: int32(s.size.x), H: int32(s.size.y)},
		&sdl.Rect{X: int32(position.x), Y: int32(position.y), W: int32(s.chunk.W), H: int32(s.chunk.H)},
		angle,
		&sdl.Point{X: int32(s.size.x / 2), Y: int32(s.size.y / 2)},
		sdl.FLIP_NONE,
	)
}
