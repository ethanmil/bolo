package maps

import (
	"image"
	"log"

	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

// Tile -
type Tile struct {
	sprite   *ebiten.Image
	typ      string
	position physics.Vector
	isHighlighted bool
}

// NewTile -
func NewTile(typ string, position physics.Vector, art *ebiten.Image) (t Tile) {
	switch typ {
	case "0": // ocean
		t.sprite = art.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
		break
	case "1": // water
		t.sprite = art.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image)
		break
	case "32": // top-bottom-road
		t.sprite = art.SubImage(image.Rect(32, 32, 64, 64)).(*ebiten.Image)
		break
	case "33": // grass
		t.sprite = art.SubImage(image.Rect(64, 32, 96, 64)).(*ebiten.Image)
		break
	case "34": // forest
		t.sprite = art.SubImage(image.Rect(96, 32, 128, 64)).(*ebiten.Image)
		break
	case "37": // single-wall
		t.sprite = art.SubImage(image.Rect(192, 32, 224, 64)).(*ebiten.Image)
		break
	default: // water
		t.sprite = art.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
	}

	t.typ = typ
	t.position = position

	return t
}

// Draw -
func (t *Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.position.X, t.position.Y)
	if t.isHighlighted{
		op.
	}
	err := screen.DrawImage(t.sprite, op)
	if err != nil {
		log.Fatal(err)
	}
}

// Highlight -
func (t *Tile) Highlight() {

}