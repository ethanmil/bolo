package tank

import (
	"image"
	"math"
	"time"

	"github.com/ethanmil/bolo/bullet"
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
	Element       *animation.Element
	speed         float64
	lastShot      time.Time
	bulletManager *bullet.Manager
}

// NewTank -
func NewTank(position physics.Vector, art *ebiten.Image, bulletManager *bullet.Manager) Tank {
	return Tank{
		Element: &animation.Element{
			Sprite:   art.SubImage(image.Rect(0, 684, 32, 716)).(*ebiten.Image),
			Position: position,
			Angle:    physics.NewAngle(float64(0)),
		},
		bulletManager: bulletManager,
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

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		t.shoot()
	}

	t.Element.Update(t.speed, delta)
}

// Draw -
func (t *Tank) Draw(screen *ebiten.Image) {
	t.Element.Draw(screen)
}

func (t *Tank) shoot() {
	if time.Since(t.lastShot) >= bulletCooldown {
		t.bulletManager.AddBullet(t.getGunPosition(), t.Element.Angle)
		t.lastShot = time.Now()
	}
}

func (t *Tank) getGunPosition() (v physics.Vector) {
	w, h := t.Element.Sprite.Size()
	v.X = t.Element.Position.X + (math.Cos(float64(t.Element.Angle)) * float64(w) / 2)
	v.Y = t.Element.Position.Y + (math.Sin(float64(t.Element.Angle)) * float64(h) / 2)
	return v
}
