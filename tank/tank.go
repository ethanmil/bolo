package tank

import (
	"image"
	"time"

	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

const (
	speed          = 0.1
	bulletSpeed    = 0.3
	bulletCooldown = time.Millisecond * 250
)

// Tank -
type Tank struct {
	Element  *animation.Element
	speed    float64
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
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		t.Element.Angle -= 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		t.Element.Angle += 0.02
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if t.speed >= 0.005 {
			t.speed -= 0.005
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if t.speed < 1 {
			t.speed += 0.01
		}
	} else {
		if t.speed >= 0.01 {
			t.speed -= 0.01
		}
	}

	// if keys[sdl.SCANCODE_SPACE] == 1 {
	// 	t.shoot()
	// }

	t.Element.Update(t.speed, delta)
}

// Draw -
func (t *Tank) Draw(screen *ebiten.Image) {
	t.Element.Draw(screen)
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
