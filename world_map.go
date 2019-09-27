package main

import "github.com/veandco/go-sdl2/sdl"

type worldMap struct {
	size  vector
	tiles [][]tile
}

func newWorldMap(size vector) (wm worldMap) {
	wm.size = size
	for x := 0; x < int(size.x); x++ {
		for y := 0; y < int(size.y); y++ {
			wm.tiles[x][y] = newTile("water")
		}
	}

	return wm
}

func (wm *worldMap) draw() {
	
}

type tile struct {
	sprite   sprite
	position vector
}

func newTile(typ string) (t tile) {
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

	return t
}
