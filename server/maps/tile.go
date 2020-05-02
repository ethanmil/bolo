package maps

import (
	"github.com/ethanmil/bolo/lib/animation"
	"github.com/ethanmil/bolo/lib/physics"
)

// Tile -
type Tile struct {
	Speed   float32
	Element animation.Element
	Typ     string
}

// NewTile -
func NewTile(typ string, position physics.Vector) (t Tile) {
	switch typ {
	case "0": // ocean
		t.Speed = 0.2
		break
	case "1": // water
		t.Speed = 0.3
		break
	case "32": // top-bottom-road
		t.Speed = 1.5
		break
	case "33": // grass
		t.Speed = 0.8
		break
	case "34": // forest
		t.Speed = 0.4
		break
	case "37": // single-wall
		t.Speed = 0.01
		t.Element.Collision = []int{1}
		break
	default: // water
		t.Speed = 0.3
	}

	t.Typ = typ
	t.Element.Position = position

	return t
}
