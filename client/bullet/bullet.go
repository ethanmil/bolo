package bullet

import (
	"image"
	"log"
	"math"

	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

const (
	speed = 2
)

// Bullet -
type Bullet struct {
	ID        int32
	Element   *animation.Element
	initPoint physics.Vector
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
		initPoint: position,
	}
}

// Update -
func (b *Bullet) Update(delta float64) {
	movement := b.Element.Angle.GetVector()
	b.Element.Position.X += movement.X * speed * delta
	b.Element.Position.Y += movement.Y * speed * delta
}

// Draw -
func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(float64(b.Element.Angle))
	op.GeoM.Translate(b.Element.Position.X, b.Element.Position.Y)
	err := screen.DrawImage(b.Element.Sprite, op)
	if err != nil {
		log.Fatal(err)
	}
}

// IsPastRange -
func (b *Bullet) IsPastRange(r float64) bool {
	first := math.Pow(b.Element.Position.X-b.initPoint.X, 2)
	second := math.Pow(b.Element.Position.Y-b.initPoint.Y, 2)
	return math.Sqrt(first+second) > r
}
