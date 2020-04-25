package tank

import (
	"image"
	"log"
	"math"
	"time"

	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	speed          = 0.1
	bulletSpeed    = 0.3
	bulletCooldown = time.Millisecond * 250
)

// Tank -
type Tank struct {
	Element  *animation.Element
	lastShot time.Time
}

// NewTank -
func NewTank(position physics.Vector, art *ebiten.Image) Tank {
	return Tank{
		Element: &animation.Element{
			Sprite:   art.SubImage(image.Rect(0, 684, 32, 716)).(*ebiten.Image),
			Position: position,
		},
	}
}

// Update -
func (t *Tank) Update(delta float64) {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		t.Element.Angle = physics.NewAngle(math.Pi)
		println("a hit")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		t.Element.Angle = physics.NewAngle(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		t.Element.Angle = physics.NewAngle(math.Pi / 2)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		t.Element.Angle = physics.NewAngle(3 * math.Pi / 2)
	}

	// if keys[sdl.SCANCODE_SPACE] == 1 {
	// 	t.shoot()
	// }

	t.Element.Update(delta)
}

// Draw -
func (t *Tank) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.Element.Position.X, t.Element.Position.Y)

	err := screen.DrawImage(t.Element.Sprite, op)
	if err != nil {
		log.Fatal(err)
	}
}

// func (t *Tank) shoot() {
// 	if time.Since(t.lastShot) >= bulletCooldown {
// 		bullet.NewBullet(t.element.Angle, bulletSpeed, t.getGunPosition())
// 		t.lastShot = time.Now()
// 	}
// }

// func (t *Tank) getGunPosition() (v physics.Vector) {
// 	v.X = (t.element.Position.X + t.element.Sprite.Size.X/2) + (math.Cos(float64(t.element.Angle)) * t.element.Sprite.Size.X / 2)
// 	v.Y = (t.element.Position.Y + t.element.Sprite.Size.Y/2) + (math.Sin(float64(t.element.Angle)) * t.element.Sprite.Size.Y / 2)
// 	return v
// }
