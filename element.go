package main

import (
	"fmt"
	"time"
)

type element struct {
	sprite         sprite
	position       vector
	velocity       vector
	angle          angle
	collisionSpace []int
	active         bool

	// log helper
	lastLogged time.Time
}

func (e *element) draw() {
	if !e.active {
		return
	}

	if e.velocity.x != 0 || e.velocity.y != 0 {
		e.angle = newAngle(e.velocity.getAngle()) // TODO: resolve this to not be so jumpy
	}

	e.sprite.draw(e.position, e.angle.getDegrees(), renderer)
}

func (e *element) update() {

}

func (e *element) print(every time.Duration) {
	if time.Since(e.lastLogged) >= every {
		e.lastLogged = time.Now()
		println(fmt.Sprintf("element: %+v", e))
	}
}
