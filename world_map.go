package main

import "github.com/veandco/go-sdl2/sdl"

const (
	tileSize = 32
)

type worldMap struct {
	size  vector
	tiles [][]tile
}

func newWorldMap(size vector, scale float64) (wm worldMap) {
	wm.size = size
	wm.tiles = make([][]tile, int(size.x))
	for x := 0; x < int(size.x); x++ {
		wm.tiles[x] = make([]tile, int(size.y))
		for y := 0; y < int(size.y); y++ {
			wm.tiles[x][y] = newTile(
				"water",
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
	for x := 0; x < int(wm.size.x); x++ {
		for y := 0; y < int(wm.size.y); y++ {
			wm.tiles[x][y].draw()
		}
	}
}

type tile struct {
	sprite   sprite
	position vector
}

func newTile(typ string, position vector) (t tile) {
	switch typ {
	case "water":
		t.sprite = sprite{
			size: vector{
				x: 32,
				y: 32,
			},
			chunk: sdl.Rect{
				X: 0,
				Y: 0,
				H: 32,
				W: 32,
			},
		}
		break
	}

	t.position = position

	return t
}

func (t *tile) draw() {
	t.sprite.draw(t.position, 0, renderer)
}
