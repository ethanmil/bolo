package tank

import (
	"image"

	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// Tank -
type Tank struct {
	ID      int32
	Element *animation.Element
}

// NewTank -
func NewTank(id int32, position physics.Vector, angle physics.Angle, art *ebiten.Image) Tank {
	return Tank{
		ID: id,
		Element: &animation.Element{
			Sprite:   art.SubImage(image.Rect(0, 684, 32, 716)).(*ebiten.Image),
			Position: position,
			Angle:    angle,
		},
	}
}

// Draw -
func (t *Tank) Draw(screen *ebiten.Image) {
	t.Element.Draw(screen)
}
