package maps

import (
	"image"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// Tile -
type Tile struct {
	Speed   float64
	Element animation.Element
	typ     string
}

// NewTile -
func NewTile(serverTile *guide.WorldMap_Tile, art *ebiten.Image) (t Tile) {
	switch serverTile.Type {
	case "0": // ocean
		t.Element.Sprite = art.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
		t.Speed = 0.2
		break
	case "1": // water
		t.Element.Sprite = art.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image)
		t.Speed = 0.3
		break
	case "32": // top-bottom-road
		t.Element.Sprite = art.SubImage(image.Rect(32, 32, 64, 64)).(*ebiten.Image)
		t.Speed = 1.5
		break
	case "33": // grass
		t.Element.Sprite = art.SubImage(image.Rect(64, 32, 96, 64)).(*ebiten.Image)
		t.Speed = 0.8
		break
	case "34": // forest
		t.Element.Sprite = art.SubImage(image.Rect(96, 32, 128, 64)).(*ebiten.Image)
		t.Speed = 0.4
		break
	case "37": // single-wall
		t.Element.Sprite = art.SubImage(image.Rect(192, 32, 224, 64)).(*ebiten.Image)
		t.Speed = 0.01
		t.Element.Collision = []int{1}
		break
	default: // water
		t.Element.Sprite = art.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
		t.Speed = 0.3
	}

	t.typ = serverTile.Type
	t.Element.Position = physics.Vector{
		X: float64(serverTile.X),
		Y: float64(serverTile.Y),
	}

	return t
}

// Draw -
func (t *Tile) Draw(screen *ebiten.Image) {
	t.Element.Draw(screen)
}
