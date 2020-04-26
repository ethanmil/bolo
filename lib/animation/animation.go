package animation

import (
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// TODO: move out to a library later

// Element -
type Element struct {
	Sprite   *ebiten.Image
	Angle    physics.Angle
	Position physics.Vector
}

// Draw -
func (e *Element) Draw(screen *ebiten.Image) {
	w, h := e.Sprite.Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(e.Angle))
	op.GeoM.Translate(e.Position.X, e.Position.Y)

	screen.DrawImage(e.Sprite, op)
}

// Update -
func (e *Element) Update(speed, delta float64) {
	movement := e.Angle.GetVector()
	e.Position.X += movement.X * speed * delta
	e.Position.Y += movement.Y * speed * delta
}
