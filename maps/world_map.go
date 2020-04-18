package maps

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ethanmil/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
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
func NewWorldMap(path string, scale float64) (wm WorldMap) {
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
			)
		}
	}

	return wm
}

// Draw -
func (wm *WorldMap) Draw(texture *sdl.Texture, renderer *sdl.Renderer) {
	for y := 0; y < int(wm.size.Y); y++ {
		for x := 0; x < int(wm.size.X); x++ {
			wm.tiles[y][x].Draw(texture, renderer)
		}
	}
}
