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

	e.sprite.draw(e.position, e.velocity.getAngle(), renderer)
}

func (e *element) update() {

}

func (e *element) print(every time.Duration) {
	if time.Since(e.lastLogged) >= every {
		e.lastLogged = time.Now()
		println(fmt.Sprintf("element: %+v", e))
	}
}
