package maps

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ethanmil/bolo/guide"
	"github.com/ethanmil/bolo/lib/physics"
)

const (
	tileSize = 32
)

// WorldMap -
type WorldMap struct {
	Size  physics.Vector
	Tiles []Tile
}

// BuildMapFromFile -
func BuildMapFromFile() *WorldMap {
	wm := &WorldMap{}
	file, err := os.Open("assets/test_map.txt")
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

	wm.Size.X = float32(len(lines[0])/2 + 1)
	wm.Size.Y = float32(len(lines))

	wm.Tiles = make([]Tile, int(wm.Size.X*wm.Size.Y))
	for y := 0; y < int(wm.Size.Y); y++ {
		for x, tileType := range strings.Split(lines[y], ",") {
			seq := x + (y * int(wm.Size.Y))
			wm.Tiles[seq] = NewTile(tileType, physics.Vector{
				X: float32(x * tileSize),
				Y: float32(y * tileSize),
			})
		}
	}

	return wm
}

// GetStateMap -
func (wm *WorldMap) GetStateMap() *guide.WorldMap {
	tiles := []string{}
	for i := range wm.Tiles {
		tiles = append(tiles, wm.Tiles[i].typ)
	}

	return &guide.WorldMap{
		SizeH: int32(wm.Size.Y),
		SizeW: int32(wm.Size.X),
		Tiles: tiles,
	}
}

// GetTileAt -
func (wm *WorldMap) GetTileAt(x, y float32) *Tile {
	seq := int(x + (y * wm.Size.Y))
	return &wm.Tiles[seq]
}

// // GetTilesWithin -
// func (wm *WorldMap) GetTilesWithin(x1, y1, x2, y2 float32) (t []Tile) {
// 	tilee := *wm.GetTileAt(x1, y1)
// 	tilee.Element.Highlight()

// 	t = []Tile{tilee}
// 	for x := x1 / tileSize; x <= x2/tileSize; x++ {
// 		for y := y1 / tileSize; y <= y2/tileSize; y++ {
// 			wm.Tiles[int(y)][int(x)].Element.Highlight()
// 			t = append(t, wm.Tiles[int(y)][int(x)])
// 		}
// 	}

// 	return t
// }
