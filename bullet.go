package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type bullet struct {
	element *element
}

var bullets = make([]*bullet, 50)

func newBullet(angle angle, speed float64, position vector) {
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
			speed:    speed,
			position: position,
		},
	}

	bullets = append(bullets, bullet)
}

func (b *bullet) update() {
	movement := b.element.angle.getVector()
	b.element.position.x += movement.x * b.element.speed * delta
	b.element.position.y += movement.y * b.element.speed * delta
}
