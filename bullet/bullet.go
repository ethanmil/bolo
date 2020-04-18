package bullet

import (
	"github.com/ethanmil/go-engine/animation"
	"github.com/ethanmil/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
)

// Bullet -
type Bullet struct {
	element *animation.Element
}

// NewBullet -
func NewBullet(angle physics.Angle, speed float64, position physics.Vector) Bullet {
	return Bullet{
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

	// bullets = append(bullets, bullet)
}

// Update -
func (b *Bullet) Update(delta float64) {
	movement := b.element.Angle.GetVector()
	b.element.Position.X += movement.X * b.element.Speed * delta
	b.element.Position.Y += movement.Y * b.element.Speed * delta
}

// Draw -
func (b *Bullet) Draw(texture *sdl.Texture, renderer *sdl.Renderer) {
	b.element.Draw(texture, renderer)
}
