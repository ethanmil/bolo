package tank

import (
	"image"
	"math"
	"time"

	"github.com/ethanmil/bolo/client/bullet"
	"github.com/ethanmil/bolo/client/maps"
	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

const (
	bulletSpeed    = 0.3
	bulletRange    = 50
	bulletCooldown = time.Millisecond * 250
)

// Tank -
type Tank struct {
	ID            int32
	Element       *animation.Element
	speed         float64
	lastShot      time.Time
	bulletManager *bullet.Manager
	worldMap      *maps.WorldMap
}

// NewTank -
func NewTank(id int32, position physics.Vector, art *ebiten.Image, worldMap *maps.WorldMap, bulletManager *bullet.Manager) Tank {
	return Tank{
		ID: id,
		Element: &animation.Element{
			Sprite:    art.SubImage(image.Rect(0, 684, 32, 716)).(*ebiten.Image),
			Position:  position,
			Angle:     physics.NewAngle(float64(0)),
			Collision: []int{1},
		},
		bulletManager: bulletManager,
		worldMap:      worldMap,
	}
}

// NewOtherTank -
func NewOtherTank(id int32, position physics.Vector, art *ebiten.Image) Tank {
	return Tank{
		ID: id,
		Element: &animation.Element{
			Sprite:    art.SubImage(image.Rect(0, 684, 32, 716)).(*ebiten.Image),
			Position:  position,
			Angle:     physics.NewAngle(float64(0)),
			Collision: []int{1},
		},
	}
}

// Update -
func (t *Tank) Update(delta float64) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		t.shoot()
	}

	t.handleMovement(delta)
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

func (t *Tank) handleMovement(delta float64) {
	// determine acceleration/max speed based on tile
	currentTile := t.worldMap.GetTileAt(t.Element.Position.X+16, t.Element.Position.Y+16) // TODO use width/height rather than hardcoding

	// get surrounding tiles (4)
	// surroundingTiles := t.worldMap.GetTilesWithin(t.Element.Position.X, t.Element.Position.Y, t.Element.Position.X+32, t.Element.Position.Y+32)

	if t.speed > currentTile.Speed {
		t.speed -= t.speed / 50
	}

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

	// handle collision
	var overrideVector *physics.Vector
	// for _, tile := range surroundingTiles {
	// 	for _, tileCol := range tile.Element.Collision {
	// 		for _, tankCol := range t.Element.Collision {
	// 			if tileCol == tankCol {
	// 				// use natural angle & 0 out the proper x or y
	// 				v := t.Element.Angle.GetVector()
	// 				println(fmt.Sprintf("Tank: %+v | Tile: %+v", t.Element.Position, tile.Element.Position))

	// 				// figure out whether the tile is in quadrant 1,2,3,4 relative to the tank,
	// 				// for each possibility, 0 out the X or Y respectively if the current vector would ruin us.

	// 				overrideVector = &v
	// 			}
	// 		}
	// 	}
	// }
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

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if t.speed < currentTile.Speed {
			t.speed += currentTile.Speed * 0.008
		}
	} else {
		if t.speed >= 0.01 {
			t.speed -= 0.01
		}
	}

	t.Element.Update(t.speed, delta, overrideVector)
}

func (t *Tank) getGunPosition() (v physics.Vector) {
	w, h := t.Element.Sprite.Size()
	v.X = t.Element.Position.X + (math.Cos(float64(t.Element.Angle)) * float64(w) / 2)
	v.Y = t.Element.Position.Y + (math.Sin(float64(t.Element.Angle)) * float64(h) / 2)
	return v
}
