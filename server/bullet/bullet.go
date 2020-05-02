package bullet

import (
	"time"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
)

const (
	speed = 0.02
)

// Bullet -
type Bullet struct {
	ID      int32
	Element *animation.Element
}

// NewBullet -
func NewBullet(id int32, position physics.Vector, angle physics.Angle) Bullet {
	return Bullet{
		ID: id,
		Element: &animation.Element{
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
func (b *Bullet) Update() {
	if b.Element.Updated.IsZero() {
		b.Element.Updated = time.Now()
	}
	delta := float32(time.Now().Sub(b.Element.Updated).Milliseconds())
	movement := b.Element.Angle.GetVector()
	b.Element.Position.X += movement.X * speed * delta
	b.Element.Position.Y += movement.Y * speed * delta
}
