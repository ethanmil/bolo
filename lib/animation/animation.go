package animation

import (
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// TODO: move out to a library later

// Element -
type Element struct {
	Sprite   *ebiten.Image
	Position physics.Vector
}

// Draw -
func (e *Element) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.Position.X, e.Position.Y)

	screen.DrawImage(e.Sprite, op)
}

// Update -
func (e *Element) Update(delta float64) {
	// movement := e.Angle.GetVector()
	// e.Position.X += movement.X * e.Speed * delta
	// e.Position.Y += movement.Y * e.Speed * delta
}
