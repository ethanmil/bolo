package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	tileSize = 32
)

type worldMap struct {
	size  vector
	tiles [][]tile
}

func newWorldMap(path string, scale float64) (wm worldMap) {
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

	wm.size = vector{
		x: float64(worldWidth),
		y: float64(worldHeight),
	}
	wm.tiles = make([][]tile, int(worldHeight))
	for y := 0; y < int(worldHeight); y++ {
		wm.tiles[y] = make([]tile, int(worldWidth))
		for x, tileType := range strings.Split(lines[y], ",") {

			wm.tiles[y][x] = newTile(
				tileType,
				vector{
					x: float64(x) * float64(tileSize) * scale,
					y: float64(y) * float64(tileSize) * scale,
				},
			)
		}
	}

	return wm
}

func (wm *worldMap) draw() {
	for y := 0; y < int(wm.size.y); y++ {
		for x := 0; x < int(wm.size.x); x++ {
			wm.tiles[y][x].draw()
		}
	}
}
