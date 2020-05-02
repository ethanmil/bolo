package tank

import (
	"image"
	"math"
	"time"

	"github.com/ethanmil/bolo/client/maps"
	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

const (
	bulletSpeed    = 0.3
	bulletCooldown = time.Millisecond * 250
)

// Tank -
type Tank struct {
	ID       int32
	Element  *animation.Element
	speed    float32
	lastShot time.Time
	worldMap *maps.WorldMap
	updated  time.Time
}

// NewTank -
func NewTank(id int32, position physics.Vector, art *ebiten.Image, worldMap *maps.WorldMap) Tank {
	return Tank{
		ID: id,
		Element: &animation.Element{
			Sprite:    art.SubImage(image.Rect(0, 684, 32, 716)).(*ebiten.Image),
			Position:  position,
			Angle:     physics.NewAngle(float32(0)),
			Collision: []int{1},
		},
		worldMap: worldMap,
	}
}

// NewOtherTank -
func NewOtherTank(id int32, position physics.Vector, art *ebiten.Image) Tank {
	return Tank{
		ID: id,
		Element: &animation.Element{
			Sprite:    art.SubImage(image.Rect(0, 684, 32, 716)).(*ebiten.Image),
			Position:  position,
			Angle:     physics.NewAngle(float32(0)),
			Collision: []int{1},
		},
	}
}

// HandleMovement -
func (t *Tank) HandleMovement(input *guide.UserInput) {
	delta := float32(time.Now().Sub(t.Element.Updated).Milliseconds())

	// determine acceleration/max speed based on tile
	currentTile := t.worldMap.GetTileAt(t.Element.Position.X+16, t.Element.Position.Y+16) // TODO use width/height rather than hardcoding

	if t.speed > currentTile.Speed {
		t.speed -= t.speed / 50
	}

	if input.left {
		t.Element.Angle -= 0.02
	}
	if input.right {
		t.Element.Angle += 0.02
	}
	if input.down {
		if t.speed >= 0.005 {
			t.speed -= 0.005
		}
	}

	// handle collision
	var overrideVector *physics.Vector
	v := t.Element.Angle.GetVector()
	if v.X < 0 {
		nextXPosition := t.Element.Position.X + v.X*t.speed*delta
		if t.worldMap.GetTileAt(nextXPosition, t.Element.Position.Y).Element.DoesCollide(t.Element.Collision...) {
			v.X = 0
			overrideVector = &v
		} else if t.worldMap.GetTileAt(nextXPosition, t.Element.Position.Y+32).Element.DoesCollide(t.Element.Collision...) {
			v.X = 0
			overrideVector = &v
		}
	}
	if v.Y > 0 {
		nextYPosition := t.Element.Position.Y + v.Y*t.speed*delta
		if t.worldMap.GetTileAt(t.Element.Position.X, nextYPosition+32).Element.DoesCollide(t.Element.Collision...) {
			v.Y = 0
			overrideVector = &v
		} else if t.worldMap.GetTileAt(t.Element.Position.X+32, nextYPosition+32).Element.DoesCollide(t.Element.Collision...) {
			v.Y = 0
			overrideVector = &v
		}
	}
	if v.X > 0 {
		nextXPosition := t.Element.Position.X + v.X*t.speed*delta
		if t.worldMap.GetTileAt(nextXPosition+32, t.Element.Position.Y).Element.DoesCollide(t.Element.Collision...) {
			v.X = 0
			overrideVector = &v
		} else if t.worldMap.GetTileAt(nextXPosition+32, t.Element.Position.Y+32).Element.DoesCollide(t.Element.Collision...) {
			v.X = 0
			overrideVector = &v
		}
	}
	if v.Y < 0 {
		nextYPosition := t.Element.Position.Y + v.Y*t.speed*delta
		if t.worldMap.GetTileAt(t.Element.Position.X, nextYPosition).Element.DoesCollide(t.Element.Collision...) {
			v.Y = 0
			overrideVector = &v
		} else if t.worldMap.GetTileAt(t.Element.Position.X+32, nextYPosition).Element.DoesCollide(t.Element.Collision...) {
			v.Y = 0
			overrideVector = &v
		}
	}

	if input.up {
		if t.speed < currentTile.Speed {
			t.speed += currentTile.Speed * 0.008
		}
	} else {
		if t.speed >= 0.01 {
			t.speed -= 0.01
		}
	}

	t.Element.Update(t.speed, overrideVector)
}

func (t *Tank) shoot() {
	if time.Since(t.lastShot) >= bulletCooldown {
		t.lastShot = time.Now()
	}
}

func (t *Tank) getGunPosition() (v physics.Vector) {
	w, h := t.Element.Sprite.Size()
	v.X = t.Element.Position.X + (float32(math.Cos(float64(t.Element.Angle))) * float32(w) / 2)
	v.Y = t.Element.Position.Y + (float32(math.Sin(float64(t.Element.Angle))) * float32(h) / 2)
	return v
}
