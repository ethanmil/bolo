package bullet

import (
	"image"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

const (
	speed = 2
)

// Bullet -
type Bullet struct {
	ID      int32
	Element *animation.Element
}

// NewBullet -
func NewBullet(id int32, position physics.Vector, angle physics.Angle, art *ebiten.Image) Bullet {
	return Bullet{
		ID: id,
		Element: &animation.Element{
			Sprite:   art.SubImage(image.Rect(16, 144, 22, 152)).(*ebiten.Image),
			Position: position,
			Angle:    angle,
		},
	}
}

// GetStateBullet -
func (b *Bullet) GetStateBullet() *guide.Bullet {
	return &guide.Bullet{
		Id:    b.ID,
		X:     b.Element.Position.X,
		Y:     b.Element.Position.Y,
		Angle: float32(b.Element.Angle),
	}
}

// Update -
func (b *Bullet) Update(delta float32) {
	movement := b.Element.Angle.GetVector()
	b.Element.Position.X += movement.X * speed * delta
	b.Element.Position.Y += movement.Y * speed * delta
}
