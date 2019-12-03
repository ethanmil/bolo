package main

import (
	"github.com/ethanmil/go-engine/animation"
	"github.com/ethanmil/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
)

type bullet struct {
	element *animation.Element
}

var bullets = make([]*bullet, 50)

func newBullet(angle physics.Angle, speed float64, position physics.Vector) {
	bullet := &bullet{
		element: &animation.Element{
			Sprite: animation.Sprite{
				Size: physics.Vector{
					X: 8,
					Y: 6,
				},
				Chunk: sdl.Rect{
					X: 16,
					Y: 144,
					H: 6,
					W: 8,
				},
			},
			Active:   true,
			Angle:    angle,
			Speed:    speed,
			Position: position,
		},
	}

	bullets = append(bullets, bullet)
}

func (b *bullet) update() {
	movement := b.element.Angle.GetVector()
	b.element.Position.X += movement.X * b.element.Speed * delta
	b.element.Position.Y += movement.Y * b.element.Speed * delta
}
