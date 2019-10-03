package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSpeed = 0.1
)

type bullet struct {
	element *element
}

var bullets = make([]*bullet, 50)

func newBullet(angle angle, velocity vector, position vector) {
	bullet := &bullet{
		element: &element{
			sprite: sprite{
				size: vector{
					x: 8,
					y: 6,
				},
				chunk: sdl.Rect{
					X: 16,
					Y: 144,
					H: 6,
					W: 8,
				},
			},
			active:   true,
			angle:    angle,
			velocity: velocity,
			position: position,
		},
	}

	bullets = append(bullets, bullet)
}

func (b *bullet) update() {
	b.element.position.x += b.element.velocity.x * delta
	b.element.position.y += b.element.velocity.y * delta
}
