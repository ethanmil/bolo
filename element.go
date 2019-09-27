package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type element struct {
	position       vector
	angle          float64
	size           vector
	collisionSpace []int
	active         bool
	chunk          sdl.Rect
	// log helper
	lastLogged time.Time
}

func (e *element) draw(renderer *sdl.Renderer) {
	if !e.active {
		return
	}

	renderer.CopyEx(
		art,
		&sdl.Rect{X: int32(e.chunk.X), Y: int32(e.chunk.Y), W: int32(e.size.x), H: int32(e.size.y)},
		&sdl.Rect{X: int32(e.position.x), Y: int32(e.position.y), W: int32(e.chunk.W), H: int32(e.chunk.H)},
		e.angle,
		&sdl.Point{X: int32(e.size.x / 2), Y: int32(e.size.y / 2)},
		sdl.FLIP_NONE,
	)
}

func (e *element) update() {

}

func (e *element) print(every time.Duration) {
	if time.Since(e.lastLogged) >= every {
		e.lastLogged = time.Now()
		println(fmt.Sprintf("element: %+v", e))
	}
}
