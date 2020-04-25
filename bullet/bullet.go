package bullet

import (
	"image"
	"log"

	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// Bullet -
type Bullet struct {
	Element *animation.Element
}

// NewBullet -
func NewBullet(position physics.Vector, art *ebiten.Image) Bullet {
	return Bullet{
		Element: &animation.Element{
			Sprite:   art.SubImage(image.Rect(16, 144, 22, 152)).(*ebiten.Image),
			Position: position,
		},
	}

	// bullets = append(bullets, bullet)
}

// Update -
// func (b *Bullet) Update(delta float64) {
// 	movement := b.element.Angle.GetVector()
// 	b.element.Position.X += movement.X * b.element.Speed * delta
// 	b.element.Position.Y += movement.Y * b.element.Speed * delta
// }

// Draw -
func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.Element.Position.X, b.Element.Position.Y)
	err := screen.DrawImage(b.Element.Sprite, op)
	if err != nil {
		log.Fatal(err)
	}
}
