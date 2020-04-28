package maps

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
func NewWorldMap(path string, scale float64, art *ebiten.Image) (wm *WorldMap) {
	wm = &WorldMap{}
	file, err := os.Open(path)
	if err != nil {
		println(fmt.Sprintf("Error: %+v", err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		println(fmt.Sprintf("Error: %+v", err))
	}

	worldWidth := len(lines[0])/2 + 1
	worldHeight := len(lines)

	wm.size = physics.Vector{
		X: float64(worldWidth),
		Y: float64(worldHeight),
	}
	wm.tiles = make([][]Tile, int(worldHeight))
	for y := 0; y < int(worldHeight); y++ {
		wm.tiles[y] = make([]Tile, int(worldWidth))
		for x, tileType := range strings.Split(lines[y], ",") {
			wm.tiles[y][x] = NewTile(
				tileType,
				physics.Vector{
					X: float64(x) * float64(tileSize) * scale,
					Y: float64(y) * float64(tileSize) * scale,
				},
				art,
			)
		}
	}

	return wm
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
func (wm *WorldMap) GetTileAt(x, y float64) *Tile {
	xIndex := int(x / tileSize)
	yIndex := int(y / tileSize)
	return &wm.tiles[yIndex][xIndex]
}

// GetTilesWithin -
func (wm *WorldMap) GetTilesWithin(x1, y1, x2, y2 float64) (t []Tile) {
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
