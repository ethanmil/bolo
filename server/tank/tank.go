package tank

import (
	"image"
	"log"
	"math"
	"time"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/ethanmil/bolo/server/maps"
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
	Name     string
	speed    float32
	lastShot time.Time
	worldMap *maps.WorldMap
	updated  time.Time
}

// NewTank -
func NewTank(id int32, worldMap *maps.WorldMap) Tank {
	return Tank{
		ID: id,
		Element: &animation.Element{
			Position:  physics.NewVector(200, 200),
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

// GetStateTank -
func (t *Tank) GetStateTank() *guide.Tank {
	return &guide.Tank{
		Id:    t.ID,
		X:     t.Element.Position.X,
		Y:     t.Element.Position.Y,
		Angle: float32(t.Element.Angle),
		Name:  t.Name,
	}
}

// HandleMovement -
func (t *Tank) HandleMovement(input *guide.UserInput) {
	if t.Element.Updated.IsZero() {
		t.Element.Updated = time.Now()
	}
	delta := float32(time.Now().Sub(t.Element.Updated).Milliseconds())

	// determine acceleration/max speed based on tile
	currentTile := t.worldMap.GetTileAt(t.Element.Position.X+16, t.Element.Position.Y+16) // TODO use width/height rather than hardcoding
	log.Printf("Current Tile %v", currentTile.Typ)
	// log.Printf("Tank position: %v", t.Element.Position)

	if t.speed > currentTile.MaxSpeed {
		t.speed -= t.speed / 50
	}

	if input.Left {
		t.Element.Angle -= 0.02
	}
	if input.Right {
		t.Element.Angle += 0.02
	}
	if input.Down {
		if t.speed >= 0.0005 {
			t.speed -= 0.0005
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

	if input.Up {
		if t.speed < currentTile.MaxSpeed {
			t.speed += currentTile.Speed * 0.0008
		}
	} else {
		if t.speed >= 0.001 {
			t.speed -= 0.001
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
