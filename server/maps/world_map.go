package maps

import (
	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/physics"
	"github.com/hajimehoshi/ebiten"
)

const (
	tileSize = 32
)

// WorldMap -
type WorldMap struct {
	size  physics.Vector
	tiles [][]Tile
}

// NewWorldMap -
func NewWorldMap(serverWM *guide.WorldMap, art *ebiten.Image) (wm *WorldMap) {
	wm = &WorldMap{
		size: physics.Vector{
			X: float32(serverWM.SizeW),
			Y: float32(serverWM.SizeH),
		},
		tiles: make([][]Tile, int(serverWM.SizeH)),
	}
	for y := 0; y < int(wm.size.Y); y++ {
		wm.tiles[y] = make([]Tile, int(wm.size.X))
		for x := 0; x < int(wm.size.X); x++ {
			wm.tiles[y][x] = NewTile(serverWM.Tiles[x+(y*int(wm.size.Y))], physics.Vector{
				X: float32(x * tileSize),
				Y: float32(y * tileSize),
			}, art)
		}
	}

	return
}

// Draw -
func (wm *WorldMap) Draw(screen *ebiten.Image) {
	for y := 0; y < int(wm.size.Y); y++ {
		for x := 0; x < int(wm.size.X); x++ {
			wm.tiles[y][x].Draw(screen)
		}
	}
}

// GetTileAt -
func (wm *WorldMap) GetTileAt(x, y float32) *Tile {
	xIndex := int(x / tileSize)
	yIndex := int(y / tileSize)
	return &wm.tiles[yIndex][xIndex]
}

// GetTilesWithin -
func (wm *WorldMap) GetTilesWithin(x1, y1, x2, y2 float32) (t []Tile) {
	tilee := *wm.GetTileAt(x1, y1)
	tilee.Element.Highlight()

	t = []Tile{tilee}
	for x := x1 / tileSize; x <= x2/tileSize; x++ {
		for y := y1 / tileSize; y <= y2/tileSize; y++ {
			wm.tiles[int(y)][int(x)].Element.Highlight()
			t = append(t, wm.tiles[int(y)][int(x)])
		}
	}

	return t
}
