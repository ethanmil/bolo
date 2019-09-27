package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type element struct {
	sprite         sprite
	position       vector
	velocity       vector
	collisionSpace []int
	active         bool
	// log helper
	lastLogged time.Time
}

func (e *element) draw(renderer *sdl.Renderer) {
	if !e.active {
		return
	}

	renderer.CopyEx(
		art,
		&sdl.Rect{X: int32(e.sprite.chunk.X), Y: int32(e.sprite.chunk.Y), W: int32(e.sprite.size.x), H: int32(e.sprite.size.y)},
		&sdl.Rect{X: int32(e.position.x), Y: int32(e.position.y), W: int32(e.sprite.chunk.W), H: int32(e.sprite.chunk.H)},
		e.velocity.getAngle(),
		&sdl.Point{X: int32(e.sprite.size.x / 2), Y: int32(e.sprite.size.y / 2)},
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
