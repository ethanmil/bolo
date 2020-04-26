package maps

import (
	"image"
	"log"

	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// Tile -
type Tile struct {
	Speed         float64
	sprite        *ebiten.Image
	typ           string
	position      physics.Vector
	isHighlighted bool
}

// NewTile -
func NewTile(typ string, position physics.Vector, art *ebiten.Image) (t Tile) {
	switch typ {
	case "0": // ocean
		t.sprite = art.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
		t.Speed = 0.2
		break
	case "1": // water
		t.sprite = art.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image)
		t.Speed = 0.3
		break
	case "32": // top-bottom-road
		t.sprite = art.SubImage(image.Rect(32, 32, 64, 64)).(*ebiten.Image)
		t.Speed = 1.5
		break
	case "33": // grass
		t.sprite = art.SubImage(image.Rect(64, 32, 96, 64)).(*ebiten.Image)
		t.Speed = 0.8
		break
	case "34": // forest
		t.sprite = art.SubImage(image.Rect(96, 32, 128, 64)).(*ebiten.Image)
		t.Speed = 0.4
		break
	case "37": // single-wall
		t.sprite = art.SubImage(image.Rect(192, 32, 224, 64)).(*ebiten.Image)
		t.Speed = 0.01
		break
	default: // water
		t.sprite = art.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
		t.Speed = 0.3
	}

	t.typ = typ
	t.position = position

	return t
}

// Draw -
func (t *Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.position.X, t.position.Y)
	if t.isHighlighted {
		op.ColorM.Scale(255, 255, 0, 0.8)
	}
	err := screen.DrawImage(t.sprite, op)
	if err != nil {
		log.Fatal(err)
	}
}

// Highlight -
func (t *Tile) Highlight() {
	t.isHighlighted = true
}
