package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	tankSpeed      = 0.1
	bulletSpeed    = 0.3
	bulletCooldown = time.Millisecond * 250
)

type tank struct {
	element  *element
	lastShot time.Time
}

func newTank() (t tank) {
	t.element = &element{
		sprite: sprite{
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
		},
		active: true,
		speed:  tankSpeed,
	}

	return t
}

func (t *tank) shoot() {
	if time.Since(t.lastShot) >= bulletCooldown {
		newBullet(t.element.angle, bulletSpeed, t.getGunPosition())
		t.lastShot = time.Now()
	}
}

func (t *tank) getGunPosition() (v vector) {
	v.x = (t.element.position.x + t.element.sprite.size.x/2) + (math.Cos(t.element.angle.radians) * t.element.sprite.size.x / 2)
	v.y = (t.element.position.y + t.element.sprite.size.y/2) + (math.Sin(t.element.angle.radians) * t.element.sprite.size.y / 2)
	return v
}

func (t *tank) update() {
	keys := sdl.GetKeyboardState()
	move := false
	if keys[sdl.SCANCODE_LEFT] == 1 {
		t.element.angle = newAngle(math.Pi)
		move = true
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		t.element.angle = newAngle(0)
		move = true
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		t.element.angle = newAngle(math.Pi / 2)
		move = true
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		t.element.angle = newAngle(3 * math.Pi / 2)
		move = true
	}

	if move {
		t.element.speed = tankSpeed
	} else {
		t.element.speed = 0
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		t.shoot()
	}

	t.element.update()
}
