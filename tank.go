package main

import (
	"math"
	"time"

	"github.com/ethanmil/go-engine/animation"
	"github.com/ethanmil/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	tankSpeed      = 0.1
	bulletSpeed    = 0.3
	bulletCooldown = time.Millisecond * 250
)

type tank struct {
	element  *animation.Element
	lastShot time.Time
}

func newTank() (t tank) {
	t.element = &animation.Element{
		Sprite: animation.Sprite{
			Size: physics.Vector{
				X: 32,
				Y: 32,
			},
			Chunk: sdl.Rect{
				X: 0,
				Y: 684,
				H: 32,
				W: 32,
			},
		},
		Active: true,
		Speed:  tankSpeed,
	}

	return t
}

func (t *tank) shoot() {
	if time.Since(t.lastShot) >= bulletCooldown {
		newBullet(t.element.Angle, bulletSpeed, t.getGunPosition())
		t.lastShot = time.Now()
	}
}

func (t *tank) getGunPosition() (v physics.Vector) {
	v.X = (t.element.Position.X + t.element.Sprite.Size.X/2) + (math.Cos(float64(t.element.Angle)) * t.element.Sprite.Size.X / 2)
	v.Y = (t.element.Position.Y + t.element.Sprite.Size.Y/2) + (math.Sin(float64(t.element.Angle)) * t.element.Sprite.Size.Y / 2)
	return v
}

func (t *tank) update() {
	keys := sdl.GetKeyboardState()
	move := false
	if keys[sdl.SCANCODE_A] == 1 {
		t.element.Angle = physics.NewAngle(math.Pi)
		move = true
	}
	if keys[sdl.SCANCODE_D] == 1 {
		t.element.Angle = physics.NewAngle(0)
		move = true
	}
	if keys[sdl.SCANCODE_S] == 1 {
		t.element.Angle = physics.NewAngle(math.Pi / 2)
		move = true
	}
	if keys[sdl.SCANCODE_W] == 1 {
		t.element.Angle = physics.NewAngle(3 * math.Pi / 2)
		move = true
	}

	if move {
		t.element.Speed = tankSpeed
	} else {
		t.element.Speed = 0
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		t.shoot()
	}

	t.element.Update(delta)
}
